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
	"time"
)

// Client 是 HTTP 客户端，用于发送请求并处理响应
type Client struct {
	BaseURL     string
	HTTPClient  *http.Client
	Logger      Logger
	CacheConfig *CacheConfig
}

// Request 发送 HTTP 请求并返回响应
func (c *Client) Request(options *RequestOptions) (*Response, error) {
	if options.Timeout == 0 {
		options.Timeout = 1000
	}
	if options.MaxContentLength == 0 {
		options.MaxContentLength = 2000
	}
	if options.MaxBodyLength == 0 {
		options.MaxBodyLength = 2000
	}
	if options.ResponseType == "" {
		options.ResponseType = "json"
	}
	if options.ResponseEncoding == "" {
		options.ResponseEncoding = "utf8"
	}
	if options.MaxRedirects == 0 {
		options.MaxRedirects = 21
	}
	if options.Method == "" {
		options.Method = MethodGet
	}
	if !options.Decompress {
		options.Decompress = true
	}

	startTime := time.Now()
	var fullURL string
	if c.BaseURL != "" {
		var err error
		fullURL, err = url.JoinPath(c.BaseURL, options.URL)
		if err != nil {
			return nil, err
		}
	} else if options.BaseURL != "" {
		var err error
		fullURL, err = url.JoinPath(options.BaseURL, options.URL)
		if err != nil {
			return nil, err
		}
	} else {
		fullURL = options.URL
	}

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

	var bodyReader io.Reader
	var bodyLength int64

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
		if options.MaxBodyLength > 0 && bodyLength > int64(options.MaxBodyLength) {
			return nil, errors.New("request body length exceeded maxBodyLength")
		}

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

	for _, interceptor := range options.InterceptorOptions.RequestInterceptors {
		err = interceptor(req)
		if err != nil {
			return nil, fmt.Errorf("request interceptor failed: %w", err)
		}
	}

	if options.Headers == nil {
		options.Headers = make(map[string]string)
	}

	if options.Body != nil {
		if _, exists := options.Headers["Content-Type"]; !exists {
			options.Headers["Content-Type"] = "application/json"
		}
	}

	// 设置请求头
	for key, value := range options.Headers {
		req.Header.Set(key, value)
	}

	if options.Auth != nil {
		auth := options.Auth.Username + ":" + options.Auth.Password
		basicAuth := base64.StdEncoding.EncodeToString([]byte(auth))
		req.Header.Set("Authorization", "Basic "+basicAuth)
	}

	if c.Logger != nil {
		c.Logger.LogRequest(req, options.LogLevel)
	}

	httpClient := &http.Client{
		Timeout: time.Duration(options.Timeout) * time.Millisecond,
	}

	if options.MaxRedirects > 0 {
		httpClient.CheckRedirect = func(_ *http.Request, via []*http.Request) error {
			if len(via) >= options.MaxRedirects {
				return fmt.Errorf("too many redirects (max: %d)", options.MaxRedirects)
			}
			return nil
		}
	}

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

	resp, err := httpClient.Do(req)
	if err != nil {
		if c.Logger != nil {
			c.Logger.LogError(err, options.LogLevel)
		}
		return nil, err
	}

	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			if err != nil {
				err = fmt.Errorf("%w; failed to close response body: %v", err, cerr)
			} else {
				err = fmt.Errorf("failed to close response body: %v", cerr)
			}
		}
	}()

	var responseBody []byte
	limitedReader := io.LimitReader(resp.Body, options.MaxContentLength+1)

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

	if int64(len(responseBody)) > options.MaxContentLength {
		return nil, errors.New("response content length exceeded maxContentLength")
	}

	duration := time.Since(startTime)

	if c.Logger != nil {
		c.Logger.LogResponse(resp, responseBody, duration, options.LogLevel)
	}

	if options.ValidateStatus != nil && !(options.ValidateStatus(resp.StatusCode)) {
		return nil, fmt.Errorf("Request failed with status code: %v", resp.StatusCode)
	}

	for _, interceptor := range options.InterceptorOptions.ResponseInterceptors {
		err = interceptor(resp)
		if err != nil {
			return nil, fmt.Errorf("response interceptor failed: %w", err)
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

	return &Response{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       responseBody,
	}, err
}

// NewClientWithCache 创建一个启用缓存的客户端
func NewClientWithCache(baseURL string, cacheConfig *CacheConfig) *Client {
	return &Client{
		BaseURL:     baseURL,
		HTTPClient:  &http.Client{},
		Logger:      NewLogger(LevelNone),
		CacheConfig: cacheConfig,
	}
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
