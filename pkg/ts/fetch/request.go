package fetch

import (
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"time"
)

// requestConfig 内部使用的请求配置
type requestConfig struct {
	method       string
	headers      http.Header
	body         []byte
	timeout      time.Duration
	proxy        string
	maxRedirects int // 0 表示禁止重定向
	jar          http.CookieJar
}

// RequestOption 用于配置 Fetch 请求
type RequestOption func(*requestConfig)

// WithMethod 设置 HTTP 方法，如 GET、POST、PUT、DELETE 等
func WithMethod(method string) RequestOption {
	return func(c *requestConfig) {
		c.method = method
	}
}

// WithHeader 添加单个请求头
func WithHeader(key, value string) RequestOption {
	return func(c *requestConfig) {
		c.headers.Set(key, value)
	}
}

// WithHeaders 批量设置请求头
func WithHeaders(headers map[string]string) RequestOption {
	return func(c *requestConfig) {
		for k, v := range headers {
			c.headers.Set(k, v)
		}
	}
}

// WithBody 设置请求体，接收字符串或字节切片
func WithBody(body []byte) RequestOption {
	return func(c *requestConfig) {
		c.body = body
	}
}

// WithJSONBody 方便地将对象序列化为 JSON 并设置 Content-Type
func WithJSONBody(v any) RequestOption {
	return func(c *requestConfig) {
		data, _ := json.Marshal(v) // 实际可用 encoding/json
		c.body = data
		c.headers.Set("Content-Type", "application/json")
	}
}

// WithTimeout 设置请求超时时间
func WithTimeout(d time.Duration) RequestOption {
	return func(c *requestConfig) {
		c.timeout = d
	}
}

// WithProxy 设置 HTTP 代理，如 "http://127.0.0.1:8080"
func WithProxy(proxyURL string) RequestOption {
	return func(c *requestConfig) {
		c.proxy = proxyURL
	}
}

// WithMaxRedirects 设置最大重定向次数，0 表示禁止重定向
func WithMaxRedirects(n int) RequestOption {
	return func(c *requestConfig) {
		c.maxRedirects = n
	}
}

// WithCookieJar 使用自定义的 cookie 容器
func WithCookieJar(jar http.CookieJar) RequestOption {
	return func(c *requestConfig) {
		c.jar = jar
	}
}

// NewCookieJar 创建一个内存 cookie jar，可跨请求共享
func NewCookieJar() http.CookieJar {
	jar, _ := cookiejar.New(nil)
	return jar
}
