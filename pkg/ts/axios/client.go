package axios

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// Client 是 HTTP 客户端，用于发送请求并处理响应
type Client struct {
	// Options 每次请求的默认配置选项
	Options *AxiosOptions
	// CacheConfig Client的缓存配置，用于自定义缓存行为，每个请求可以共用缓存
	CacheConfig *CacheConfig
	// logger 是日志记录器，用于记录请求和响应，每个请求使用同一个日志记录器
	logger Logger
	// httpClient 是 HTTP 客户端，用于发送请求，每个请求使用同一个 HTTP 客户端
	httpClient *http.Client
	// 请求队列，用于处理并发请求
	queue     sync.WaitGroup
	queueChan chan struct{}
}

// New
func New() *Client {
	return &Client{
		Options: &AxiosOptions{
			Timeout:          1000,
			ResponseType:     "json",
			ResponseEncoding: "utf8",
			MaxContentLength: 1024 * 1024, // 1MB
			MaxBodyLength:    2000,
			MaxRedirects:     21,
			Decompress:       true,
			ValidateStatus:   nil,
			Headers: map[string]string{
				"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36 Edg/133.0.0.0",
				"Accept-Encoding": "gzip",
				"Accept-Language": "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3",
				"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
				"Content-Type":    "application/json",
				"Connection":      "keep-alive",
			},
		},
		httpClient: &http.Client{},
		logger:     NewLogger(LevelError),
		queueChan:  make(chan struct{}, 64), //
	}
}

// NewClient 创建一个新的 HTTP 客户端
func NewClient(baseURL string) *Client {
	client := New()
	if baseURL != "" {
		client.Options.BaseURL = baseURL
	}
	return client
}

// NewClientWithCache 创建一个启用缓存的客户端
func NewClientWithCache(baseURL string, cacheConfig *CacheConfig) *Client {
	client := NewClient(baseURL)
	client.CacheConfig = cacheConfig
	return client
}

// Request 发送 HTTP 请求并返回响应
func (c *Client) Request(configs ...*AxiosOptions) (*Response, error) {
	options := new(*c.Options)
	for _, opt := range configs {
		mergeOptions(options, opt)
	}
	// 构建完整的 URL
	fullURL := options.URL
	if options.BaseURL != "" {
		var err error
		fullURL, err = url.JoinPath(options.BaseURL, fullURL)
		if err != nil {
			return nil, err
		}
	}
	// 处理查询参数
	if len(options.Params) > 0 {
		parsedURL, err := url.Parse(fullURL)
		if err != nil {
			return nil, err
		}
		q := parsedURL.Query()
		for k, v := range options.Params {
			q.Add(k, v)
		}
		parsedURL.RawQuery = q.Encode()
		fullURL = parsedURL.String()
	}

	// 缓存检查：在发起请求前尝试获取缓存的响应
	var cacheKey string
	shouldCache := shouldCacheRequest(c.CacheConfig, options)

	if shouldCache && !shouldForceRefresh(options) {
		cacheKey = generateCacheKey(c.CacheConfig, options, fullURL)
		if cachedEntry := c.CacheConfig.Cache.Get(cacheKey); cachedEntry != nil {
			// 返回缓存的响应
			return &Response{
				StatusCode: cachedEntry.StatusCode,
				Headers:    cachedEntry.Headers,
				Body:       cachedEntry.Body,
			}, nil
		}
	}
	//
	var bodyReader io.Reader
	var bodyLength int64
	// 处理请求体
	if options.Body != nil {
		switch v := options.Body.(type) {
		case string:
			bodyReader = strings.NewReader(v)
			bodyLength = int64(len(v))
		case []byte:
			bodyReader = bytes.NewReader(v)
			bodyLength = int64(len(v))
		default:
			jsonBody, err := json.Marshal(options.Body)
			if err != nil {
				return nil, err
			}
			bodyReader = bytes.NewBuffer(jsonBody)
			bodyLength = int64(len(jsonBody))
		}
		// 检查请求体长度是否超过最大限制
		if options.MaxBodyLength > 0 && bodyLength > int64(options.MaxBodyLength) {
			return nil, errors.New("请求体内容长度超过最大限制")
		}
		// 处理上传进度回调
		if options.Body != nil && options.OnUploadProgress != nil {
			bodyReader = &ProgressReader{
				reader:     bodyReader,
				total:      bodyLength,
				onProgress: options.OnUploadProgress,
			}
		}
	}

	// 创建请求
	req, err := http.NewRequest(string(options.Method), fullURL, bodyReader)
	if err != nil {
		return nil, err
	}
	// 执行请求拦截器
	for _, interceptor := range options.InterceptorOptions.RequestInterceptors {
		err = interceptor(req)
		if err != nil {
			return nil, fmt.Errorf("请求拦截器失败: %v, %w", err, err)
		}
	}

	// 设置请求头
	for key, value := range options.Headers {
		req.Header.Set(key, value)
	}

	// 处理认证
	if options.Auth != nil {
		auth := options.Auth.Username + ":" + options.Auth.Password
		basicAuth := base64.StdEncoding.EncodeToString([]byte(auth))
		req.Header.Set("Authorization", "Basic "+basicAuth)
	}
	// 记录日志
	if c.logger != nil {
		c.logger.LogRequest(req)
	}

	httpClient := c.httpClient
	httpClient.Timeout = time.Duration(options.Timeout) * time.Millisecond
	// 处理重定向
	if options.MaxRedirects > 0 {
		httpClient.CheckRedirect = func(_ *http.Request, via []*http.Request) error {
			if len(via) >= options.MaxRedirects {
				return fmt.Errorf("重定向次数超过最大重定向次数: %d", options.MaxRedirects)
			}
			return nil
		}
	}
	// 处理代理
	if options.Proxy != nil {
		proxyStr := fmt.Sprintf("%s://%s:%d", options.Proxy.Protocol, options.Proxy.Host, options.Proxy.Port)
		proxyURL, err := url.Parse(proxyStr)
		if err != nil {
			return nil, err
		}
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
		if options.Proxy.Auth != nil {
			auth := options.Proxy.Auth.Username + ":" + options.Proxy.Auth.Password
			basicAuth := base64.StdEncoding.EncodeToString([]byte(auth))
			transport.ProxyConnectHeader = http.Header{
				"Proxy-Authorization": {"Basic " + basicAuth},
			}
		}
		httpClient.Transport = transport
	}
	// 记录请求时间
	startTime := time.Now()
	// 发起请求
	resp, err := httpClient.Do(req)
	if err != nil {
		if c.logger != nil {
			c.logger.LogError(err)
		}
		return nil, err
	}

	// 关闭响应体
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			if err != nil {
				err = fmt.Errorf("%w; 关闭响应体失败，错误: %v", err, cerr)
			} else {
				err = fmt.Errorf("关闭响应体失败，错误: %v", cerr)
			}
		}
	}()

	var responseBody []byte
	limitedReader := io.LimitReader(resp.Body, options.MaxContentLength+1)
	// 处理下载进度回调
	if options.OnDownloadProgress != nil {
		buf := &bytes.Buffer{}
		progressWriter := &ProgressWriter{
			writer:     buf,
			total:      resp.ContentLength,
			onProgress: options.OnDownloadProgress,
		}
		_, err = io.Copy(progressWriter, limitedReader)
		if err != nil {
			return nil, err
		}
		responseBody = buf.Bytes()
	} else {
		responseBody, err = io.ReadAll(limitedReader)
		if err != nil {
			return nil, err
		}
	}
	// 检查响应内容长度是否超过最大限制
	if int64(len(responseBody)) > options.MaxContentLength {
		return nil, errors.New("响应内容长度超过最大限制")
	}
	// 记录响应时间
	duration := time.Since(startTime)
	if c.logger != nil {
		// 相应时间太长时，发送警告日志
		if duration > time.Minute {

		}
		c.logger.LogResponse(resp, responseBody, duration)
	}

	// 验证响应状态码
	if options.ValidateStatus != nil && !(options.ValidateStatus(resp.StatusCode)) {
		return nil, fmt.Errorf("请求失败，状态码: %v", resp.StatusCode)
	}
	// 执行响应拦截器
	for _, interceptor := range options.InterceptorOptions.ResponseInterceptors {
		err = interceptor(resp)
		if err != nil {
			return nil, fmt.Errorf("响应拦截器失败，错误: %w", err)
		}
	}

	// 缓存存储：将成功的响应保存到缓存中
	if shouldCache && resp.StatusCode >= 200 && resp.StatusCode < 300 {
		ttl := getCacheTTL(c.CacheConfig, options)
		if ttl > 0 {
			if cacheKey == "" {
				cacheKey = generateCacheKey(c.CacheConfig, options, fullURL)
			}
			c.CacheConfig.Cache.Set(cacheKey, &CacheEntry{
				Body:       responseBody,
				StatusCode: resp.StatusCode,
				Headers:    resp.Header.Clone(),
				CreatedAt:  time.Now(),
			}, ttl)
		}
	}
	// 返回响应
	return &Response{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       responseBody,
	}, err
}

// RequestAsync 异步发送 HTTP 请求，返回 Promise
func (c *Client) RequestAsync(options *AxiosOptions) *Promise {
	promise := NewPromise()
	// 同步模式
	if c.queueChan == nil {
		resp, err := c.Request(options)
		promise.resolve(resp, err)
	} else {
		// 异步模式
		c.queue.Go(func() {
			// 超过缓冲区大小时会自动阻塞
			c.queueChan <- struct{}{}
			resp, err := c.Request(options)
			promise.resolve(resp, err)
			<-c.queueChan
		})
	}
	return promise
}

// SetMaxPendingRequests 设置最大并发请求数
// 默认值为 64
func (c *Client) SetMaxPendingRequests(max int) {
	if max <= 0 {
		max = 64
	}
	c.queueChan = make(chan struct{}, max)
}

// Wait 等待所有队列中的请求完成
// 用于在应用退出前确保所有请求完成，避免资源泄漏
// 一般不需要调用
func (c *Client) Wait() {
	c.queue.Wait()
}

// SetCache 设置客户端的缓存配置
func (c *Client) SetCache(config *CacheConfig) {
	c.CacheConfig = config
}

// ClearCache 清除客户端缓存中的所有条目
func (c *Client) ClearCache() {
	if c.CacheConfig != nil && c.CacheConfig.Cache != nil {
		c.CacheConfig.Cache.Clear()
	}
}

// CacheStats 返回客户端的缓存统计信息
func (c *Client) CacheStats() *CacheStats {
	if c.CacheConfig != nil && c.CacheConfig.Cache != nil {
		stats := c.CacheConfig.Cache.Stats()
		return &stats
	}
	return nil
}
