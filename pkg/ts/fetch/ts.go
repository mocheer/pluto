package fetch

import (
	"bytes"
	"net/http"
	"time"

	"github.com/gocolly/colly/v2"
)

type FetchClient struct {
	c              *colly.Collector
	MaxConcurrency int                 // 最大并发请求数量
	queue          chan *requestConfig // 任务队列，用于异步处理
}

// fetch 发起 HTTP 请求，返回标准化的响应或错误。
func (f *FetchClient) Fetch(u string, opts ...RequestOption) (*Response, error) {
	// colly.Async(true)
	// 默认配置
	cfg := &requestConfig{
		method:       http.MethodGet,
		headers:      http.Header{},
		timeout:      30 * time.Second,
		maxRedirects: 5,
	}
	// 应用用户选项
	for _, opt := range opts {
		opt(cfg)
	}

	c := f.c
	// 配置代理
	if cfg.proxy != "" {
		err := c.SetProxy(cfg.proxy)
		if err != nil {
			return nil, err
		}
	}

	// 配置重定向策略
	c.SetRedirectHandler(func(req *http.Request, via []*http.Request) error {
		if cfg.maxRedirects == 0 {
			return http.ErrUseLastResponse // 不跟随重定向
		}
		if len(via) >= cfg.maxRedirects {
			return http.ErrUseLastResponse
		}
		// 默认跟随，但要注意，colly 会保留 cookie，这符合预期
		return nil
	})

	// 如果用户提供了 Cookie Jar，则替换 colly 的默认 Jar
	if cfg.jar != nil {
		c.SetCookieJar(cfg.jar)
	}

	// 构造 *http.Request
	// 我们使用 colly.Request 来完成，但底层的 net/http 请求需要手动注入 body 等。
	// colly 的低层请求由 colly.Request 管理，这里我们可以直接使用 c.Request。

	// 准备响应容器
	res := Response{}
	res.Headers = http.Header{}

	//
	c.OnRequest(func(r *colly.Request) {
		host := r.URL.Hostname()
		r.Headers.Set("Host", host)
		r.Headers.Set("Origin", host)
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*") //
		r.Headers.Set("Referer", r.URL.String())
		//关键头，如果没有,大概率会返回错误
		r.Headers.Set("Accept-Encoding", "gzip, deflate, br, zstd")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
		//Do NotTrack(DNT)实际是一个涉及隐私保护的协议,它是用户和网站之间的一个“君子协定”,通过此协定,用户可以允许也可以禁止网站搜集自己在网上的隐私踪迹
		r.Headers.Set("DNT", "1")

	})

	// 注册回调，捕获响应
	c.OnResponse(func(r *colly.Response) {
		res.StatusCode = r.StatusCode
		res.Status = http.StatusText(r.StatusCode)
		// 复制 headers
		for key, values := range *r.Headers {
			for _, v := range values {
				res.Headers.Add(key, v)
			}
		}
		// 复制 body
		res.Body = make([]byte, len(r.Body))
		copy(res.Body, r.Body)
		res.RequestURL = r.Request.URL.String()
	})

	// 错误处理（网络错误、超时等）
	c.OnError(func(r *colly.Response, err error) {
		// 如果已经收到响应（如4xx/5xx），则保留响应数据，并清空错误
		if r != nil {
			res.StatusCode = r.StatusCode
			res.Status = http.StatusText(r.StatusCode)
			if r.Headers != nil {
				for key, values := range *r.Headers {
					for _, v := range values {
						res.Headers.Add(key, v)
					}
				}
			}
			res.Body = make([]byte, len(r.Body))
			copy(res.Body, r.Body)
			res.RequestURL = r.Request.URL.String()
			// 将错误挂到响应上，由调用者决定
			res.err = err
		} else {
			// 完全没有响应，直接设置错误
			res.err = err
		}
	})
	collyCtx := colly.NewContext()
	err := c.Request(cfg.method, u, bytes.NewReader(cfg.body), collyCtx, cfg.headers)
	if err != nil {
		return nil, err
	}

	// 等待请求完成 (colly/v2版本的Request是同步的，会阻塞)
	// 检查是否有错误记录在响应中
	if res.err != nil {
		return &res, res.err
	}

	return &res, nil
}

// FetchAsync 异步发起 HTTP 请求，用于批量请求。
// TODO
func FetchAsync(u string, opts ...RequestOption) (*Response, error) {
	return New().Fetch(u, opts...)
}
