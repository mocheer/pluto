package axios

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sync"
	"time"
)

// Client 是 HTTP 客户端，用于发送请求并处理响应
type Client struct {
	// Options 每次请求的默认配置选项
	Options *AxiosOptions
	// CacheConfig Client的缓存配置，用于自定义缓存行为，每个请求可以共用缓存
	CacheConfig *CacheConfig
	// jar 是 CookieJar（内存存储），用于存储和管理 Cookie
	jar *cookiejar.Jar
	// logger 是日志记录器，用于记录请求和响应，每个请求使用同一个日志记录器
	logger Logger
	// 请求队列，用于处理并发请求
	queue     sync.WaitGroup
	queueChan chan struct{}
}

// New
func New() *Client {
	// 创建 CookieJar（内存存储）
	// 即使不用同一个http.Client,只要用jar创建http.client,每个请求的cookie都会被保存到jar中
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	return &Client{
		jar:       jar,
		logger:    NewLogger(LevelError),
		queueChan: make(chan struct{}, 64),
		// 默认配置
		Options: &AxiosOptions{
			Timeout:                time.Hour,     // 1小时超时,如果用来下载文件，超时时间不宜过短
			MaxResponseContentSize: math.MaxInt64, // 不限制返回的内容大小，因为下载文件可能很大，但太大，会导致内存溢出
			MaxRequestBodySize:     4096,          // 4KB
			MaxRedirects:           21,            // 最大重定向次数，默认 21 次
			ValidateStatus:         nil,           // 自定义状态码验证函数，默认 nil
			// 浏览器/Go都会自动填充Host（HTTP/1.1协议）、Connection、Content-Length，所以一般这些不需要手动设置
			// Connection: 在net/http中默认是keep-alive会自动设置,HTTP/2 协议明确禁止使用 Connection 头部
			Header: Header{
				// 默认值为 Go-http-client/1.1
				"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36 Edg/133.0.0.0",
				// 目前只支持 gzip 编码，其他编码需要根据实际情况调整
				// 浏览器支持：gzip, deflate, br
				"Accept-Encoding": "gzip",
				// 告诉服务器用户希望页面或资源返回哪种自然语言
				// 没有这个头，服务器通常会返回默认语言（常为英语）或根据 IP 猜测
				// 例如，如果用户在中国，服务器可能会返回 zh-CN 或 zh 等语言，如果IP和用户语言不一致，就会显得可疑
				"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
				// 默认值为 */*
				// img标签：image/avif,image/webp,image/apng,image/*,*/*;q=0.8
				// 导航请求：text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8
				"Accept": "*/*",
				// Content-Type 默认值为 application/json
				"Content-Type": "application/json",
				// Cache-Control 用于控制缓存行为，浏览器强制刷新时，会发送 Cache-Control: no-cache。
				"Cache-Control": "no-cache",
			},

			// 请求拦截器配置
			// 经常要根据实际情况调整请求头，比如设置 Host、Origin、Referer 等，因为很多站点有严格的请求头要求
			// 例如，一些 API 只允许从特定的域名或 IP 地址访问，而其他站点则要求请求头中包含特定的字段。
			RequestInterceptors: []RequestInterceptor{
				func(req *http.Request) error {
					// 这里采用默认的 Host 和 Origin 头,Host一般会自动填充，这里其实不需要设置
					// 实际上有些服务需要自定义的 Host 和 Origin 头，这里根据实际情况调整
					host := req.URL.Host
					// 无法通过 req.Header.Set("Host", "...") 来修改它
					// request.Header.Set("Host", host)
					// request.Host = host //只能通过request.Host来设置Host

					// Origin 头是浏览器在发起跨站请求或某些非简单请求时自动附加的，它只包含请求的来源（协议、域名和端口），不包含路径。
					// 用途：CORS (跨域资源共享) 和 安全防护（类似于 Referer 头） 用来防范 CSRF 攻击、资源盗链等。部分网站的 API 会严格校验 Origin，要求它必须与目标站点同源或出现在白名单里。
					// Origin 只到根路径。
					// 仅在特定情况下发送，如跨域请求、同域POST、PUT、DELETE 等。导航请求、同域的GET、HEAD 等不会发送 Origin 头。
					req.Header.Set("Origin", host)
					// Referer 头记录了请求的来源页面，是浏览器行为的另一重要特征，一般情况下为当前页面的 URL
					// Referer 通常包含完整的路径和参数，比 Origin 头更详细，隐私敏感的场景下，浏览器可能会仅发送 Origin。
					// 1. 首请求无 Referer 头，其他请求需要设置为上一个请求的 URL，所以这里其实要判断是否是首请求
					// 2. 从一个页面跳转到另一个页面，或页面引用了跨域资源时，会自动携带。
					// 3. 重定向时，Referer 一般需要设置为上一个请求的 URL
					// 浏览器有时会出于隐私考虑隐藏或裁剪 Referer（如使用 Referrer-Policy）,或者限制 Referer 的内容（如只发送源信息、降级等）。
					// 对爬虫来说，我们更多是反其道而行之，确保它总是携带一个合理的来源。
					req.Header.Set("Referer", req.URL.String())

					// Sec-Fetch-* 系列头用于描述请求的类型，如导航请求、资源请求、跨站请求等。
					// 以 Sec- 或 Proxy- 开头的头，都是浏览器在发起请求时自动添加的，用于描述请求的类型、模式、站点等，js端不能手动设置。
					// 1. Sec-Fetch-Dest: 描述请求的目标类型，如 document、image、script 等。
					// 2. Sec-Fetch-Mode: 描述请求的模式，如 navigate（导航）、cors（跨域）、no-cors（不跨域）、same-origin（同源）等。
					// 3. Sec-Fetch-Site: 描述请求的站点，如 same-origin、cross-origin 等。
					// 4. Sec-Fetch-User: 描述这个请求是不是由用户的主动操作（如点击、按键）直接触发的，只有两个值: ?1或者没有这个请求头
					req.Header.Set("Sec-Fetch-Dest", "document")
					req.Header.Set("Sec-Fetch-Mode", "same-origin")
					req.Header.Set("Sec-Fetch-Site", "same-origin")
					req.Header.Set("Sec-Fetch-User", "?1")
					// Do NotTrack(DNT) 这是一个被弃用的隐私保护机制
					// 这是浏览器向服务器表达用户“不希望被追踪”意愿的一种方式，但这个设计存在根本性缺陷，导致其在现实中收效甚微
					// 通过此协定,用户可以允许也可以禁止网站搜集自己在网上的隐私踪迹
					// 1 用户不希望被追踪
					// 0 用户同意被追踪
					// 这个已经被弃用，发送反而可能导致风险
					// req.Header.Set("DNT", "1")
					return nil
				},
			},
		},
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
	parsedURL, err := url.Parse(fullURL)

	if err != nil {
		return nil, err
	}
	// 处理查询参数
	if len(options.Params) > 0 {
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
	var gzipReader io.Reader
	// 处理请求体
	if options.Body != nil {
		var bodyLength int64
		gzipReader, bodyLength = createReqBody(options.Body)
		if gzipReader == nil {
			return nil, errors.New("请求体内容读取失败")
		}
		// 检查请求体长度是否超过最大限制
		if options.MaxRequestBodySize > 0 && bodyLength > int64(options.MaxRequestBodySize) {
			return nil, errors.New("请求体内容长度超过最大限制")
		}
		// 处理上传进度回调
		if options.OnUploadProgress != nil {
			gzipReader = &ProgressReader{
				reader:     gzipReader,
				total:      bodyLength,
				onProgress: options.OnUploadProgress,
			}
		}
	}

	// 创建请求
	req, err := http.NewRequest(string(options.Method), fullURL, gzipReader)

	if err != nil {
		return nil, err
	}
	// 设置认证头
	options.Header.SetAuth(options.Auth)
	// 设置请求头
	for key, value := range options.Header {
		req.Header.Set(key, value)
	}
	// 执行请求拦截器
	for _, interceptor := range options.RequestInterceptors {
		err = interceptor(req)
		if err != nil {
			return nil, fmt.Errorf("请求拦截器失败: %v, %w", err, err)
		}
	}

	// 记录日志
	if c.logger != nil {
		c.logger.LogRequest(req)
	}
	// 处理 Cookie
	if len(options.Cookies) > 0 {
		c.jar.SetCookies(parsedURL, options.Cookies)
	}
	// httpClient 是 HTTP 客户端，用于发送请求，这里每个请求都创建一个新的客户端，避免线程安全问题
	// http.Client 在发送请求前，会调用 Jar 的 Cookies() 方法获取匹配的 Cookie 并填入请求头；收到响应后调用 SetCookies() 更新存储。
	// 所有使用默认 Transport 的 Client 都会共享同一个连接池，即使每次请求都新建 Client，TCP 连接依然能被复用，所以网络延迟和吞吐量基本没有损失。
	// 1. 只有自定义Transport时，才会创建新的连接池。=> 不太影响性能
	// 2. 每次新建 Client 手动传递相同的 Jar 实例，不如直接复用 Client 简洁。 => 不太影响性能
	// http.Get 等全局函数 其实是共享 DefaultClient，复用能减少 GC 压力，有益无害，但这里因为需要并发安全，暂时不考虑复用
	httpClient := &http.Client{
		Timeout: options.Timeout,
		Jar:     c.jar,
	}
	// 处理重定向
	if options.MaxRedirects > 0 {
		httpClient.CheckRedirect = func(_ *http.Request, via []*http.Request) error {
			// 记录每次重定向的 URL
			// if c.logger != nil {
			// 	c.logger.LogRedirect(via[len(via)-1].URL.String())
			// }
			//
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
	fmt.Println(string(options.Method), fullURL, gzipReader, err, startTime)
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

	// 如果请求成功，且响应体压缩为 gzip，解压缩响应体
	if !options.Uncompressed && resp.StatusCode >= 200 && resp.StatusCode < 300 && resp.Header.Get("Content-Encoding") == "gzip" {
		bodyReader := resp.Body
		defer bodyReader.Close()
		gzipReader, err = gzip.NewReader(bodyReader)
		if err != nil {
			return nil, fmt.Errorf("解压缩响应体失败，错误: %w", err)
		}
		reader := gzipReader.(*gzip.Reader)
		// 关闭 gzip 读取器
		// defer reader.Close()
		// 这里不用关闭，因为resp.Body 已经在上面做了自动关闭
		resp.Body = reader // 替换响应体为解压缩后的读取器
		// body, err := io.ReadAll(bodyReader)
		// if err != nil {
		// 	return nil, fmt.Errorf("读取解压缩后的响应体失败，错误: %w", err)
		// }
	}
	var responseBody []byte
	maxResponseContentSize := options.MaxResponseContentSize
	// 这里有溢出风险，当options.MaxResponseContentSize=math.MaxInt64时如果再+1会变成负数，会导致读取失败：EOF
	// 所以这里需要判断是否是 math.MaxInt64，如果是，就不需要再+1了
	if maxResponseContentSize < math.MaxInt64 {
		maxResponseContentSize++
	}
	// 这里的resp.Body可能是原始响应体读取器，也可能是gzip解压缩后的读取器
	limitedReader := io.LimitReader(resp.Body, maxResponseContentSize)

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
		// resp.Body 被消费殆尽，指针到 EOF 位置
		responseBody, err = io.ReadAll(limitedReader)
		if err != nil {
			return nil, err
		}
	}
	// 检查响应内容长度是否超过最大限制
	if int64(len(responseBody)) > options.MaxResponseContentSize {
		return nil, errors.New("响应内容长度超过最大限制")
	}
	// 记录响应时间
	duration := time.Since(startTime)
	if c.logger != nil {
		c.logger.LogResponse(resp, responseBody, duration)
	}

	// 验证响应状态码
	if options.ValidateStatus != nil && !(options.ValidateStatus(resp.StatusCode)) {
		return nil, fmt.Errorf("请求失败，状态码: %v", resp.StatusCode)
	}
	// 执行响应拦截器
	for _, interceptor := range options.ResponseInterceptors {
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
	res := &Response{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       responseBody,
	}

	// 返回响应
	return res, err
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

// AddRequestInterceptor 添加请求拦截器
func (c *Client) AddRequestInterceptor(interceptor RequestInterceptor) {
	c.Options.RequestInterceptors = append(c.Options.RequestInterceptors, interceptor)
}

// AddResponseInterceptor 添加响应拦截器
func (c *Client) AddResponseInterceptor(interceptor ResponseInterceptor) {
	c.Options.ResponseInterceptors = append(c.Options.ResponseInterceptors, interceptor)
}
