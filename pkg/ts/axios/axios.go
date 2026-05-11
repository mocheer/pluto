package axios

import (
	"net/http"
)

// defaultClient 是默认的 HTTP 客户端实例
var defaultClient = &Client{HTTPClient: &http.Client{}, Logger: NewLogger(LevelNone)}

// Get 发送 HTTP GET 请求
func Get(urlStr string, options ...*RequestOptions) (*Response, error) {
	return Request("GET", urlStr, options...)
}

// GetAsync 异步发送 HTTP GET 请求，返回 Promise
func GetAsync(urlStr string, options ...*RequestOptions) *Promise {
	promise := NewPromise()

	go func() {
		resp, err := Request("GET", urlStr, options...)
		promise.resolve(resp, err)
	}()

	return promise
}

// Post 发送 HTTP POST 请求，包含请求体
func Post(urlStr string, body interface{}, options ...*RequestOptions) (*Response, error) {
	mergedOptions := mergeBodyIntoOptions(body, options)
	return Request("POST", urlStr, mergedOptions)
}

// PostAsync 异步发送 HTTP POST 请求，返回 Promise
func PostAsync(urlStr string, body interface{}, options ...*RequestOptions) *Promise {
	mergedOptions := mergeBodyIntoOptions(body, options)
	promise := NewPromise()

	go func() {
		resp, err := Request("POST", urlStr, mergedOptions)
		promise.resolve(resp, err)
	}()

	return promise
}

// mergeBodyIntoOptions 将请求体合并到请求选项中
func mergeBodyIntoOptions(body interface{}, options []*RequestOptions) *RequestOptions {
	mergedOption := &RequestOptions{
		Body: body,
	}

	if len(options) > 0 {
		*mergedOption = *options[0]
		mergedOption.Body = body
	}

	return mergedOption
}

// Put 发送 HTTP PUT 请求
func Put(urlStr string, body interface{}, options ...*RequestOptions) (*Response, error) {
	mergedOptions := mergeBodyIntoOptions(body, options)
	return Request("PUT", urlStr, mergedOptions)
}

// PutAsync 异步发送 HTTP PUT 请求，返回 Promise
func PutAsync(urlStr string, body interface{}, options ...*RequestOptions) *Promise {
	mergedOptions := mergeBodyIntoOptions(body, options)
	promise := NewPromise()

	go func() {
		resp, err := Request("PUT", urlStr, mergedOptions)
		promise.resolve(resp, err)
	}()

	return promise
}

// Delete 发送 HTTP DELETE 请求
func Delete(urlStr string, options ...*RequestOptions) (*Response, error) {
	return Request("DELETE", urlStr, options...)
}

// DeleteAsync 异步发送 HTTP DELETE 请求，返回 Promise
func DeleteAsync(urlStr string, options ...*RequestOptions) *Promise {
	promise := NewPromise()
	go func() {
		resp, err := Request("DELETE", urlStr, options...)
		promise.resolve(resp, err)
	}()
	return promise
}

// Head 发送 HTTP HEAD 请求
func Head(urlStr string, options ...*RequestOptions) (*Response, error) {
	return Request("HEAD", urlStr, options...)
}

// HeadAsync 异步发送 HTTP HEAD 请求，返回 Promise
func HeadAsync(urlStr string, options ...*RequestOptions) *Promise {
	promise := NewPromise()
	go func() {
		resp, err := Request("HEAD", urlStr, options...)
		promise.resolve(resp, err)
	}()
	return promise
}

// Options 发送 HTTP OPTIONS 请求
func Options(urlStr string, options ...*RequestOptions) (*Response, error) {
	return Request("OPTIONS", urlStr, options...)
}

// OptionsAsync 异步发送 HTTP OPTIONS 请求，返回 Promise
func OptionsAsync(urlStr string, options ...*RequestOptions) *Promise {
	promise := NewPromise()
	go func() {
		resp, err := Request("OPTIONS", urlStr, options...)
		promise.resolve(resp, err)
	}()
	return promise
}

// Patch 发送 HTTP PATCH 请求
func Patch(urlStr string, body interface{}, options ...*RequestOptions) (*Response, error) {
	mergedOptions := mergeBodyIntoOptions(body, options)
	return Request("PATCH", urlStr, mergedOptions)
}

// PatchAsync 异步发送 HTTP PATCH 请求，返回 Promise
func PatchAsync(urlStr string, body interface{}, options ...*RequestOptions) *Promise {
	mergedOptions := mergeBodyIntoOptions(body, options)
	promise := NewPromise()

	go func() {
		resp, err := Request("PATCH", urlStr, mergedOptions)
		promise.resolve(resp, err)
	}()

	return promise
}

// Request 发送 HTTP 请求，使用默认客户端
func Request(method MethodType, urlStr string, options ...*RequestOptions) (*Response, error) {
	reqOptions := &RequestOptions{
		Method:           "GET",
		URL:              urlStr,
		Timeout:          1000,
		ResponseType:     "json",
		ResponseEncoding: "utf8",
		MaxContentLength: 2000,
		MaxBodyLength:    2000,
		MaxRedirects:     21,
		Decompress:       true,
		ValidateStatus:   nil,
	}

	if len(options) > 0 && options[0] != nil {
		mergeOptions(reqOptions, options[0])
	}

	if method != "" {
		reqOptions.Method = method
	}

	return defaultClient.Request(reqOptions)
}

// RequestAsync 异步发送 HTTP 请求，返回 Promise
func RequestAsync(method MethodType, urlStr string, options ...*RequestOptions) *Promise {
	resp, err := Request(method, urlStr, options...)
	return &Promise{response: resp, err: err}
}

// SetBaseURL 设置默认客户端的 BaseURL
func SetBaseURL(baseURL string) {
	defaultClient.BaseURL = baseURL
}

// NewClient 创建一个新的 HTTP 客户端
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
		Logger:     NewLogger(LevelNone),
	}
}

// mergeOptions 将源请求选项合并到目标请求选项中
func mergeOptions(dst, src *RequestOptions) {
	if src.Method != "" {
		dst.Method = src.Method
	}
	if src.URL != "" {
		dst.URL = src.URL
	}
	if src.BaseURL != "" {
		dst.BaseURL = src.BaseURL
	}
	if src.Params != nil {
		dst.Params = src.Params
	}
	if src.Body != nil {
		dst.Body = src.Body
	}
	if src.Headers != nil {
		dst.Headers = src.Headers
	}
	if src.Timeout != 0 {
		dst.Timeout = src.Timeout
	}
	if src.Auth != nil {
		dst.Auth = src.Auth
	}
	if src.ResponseType != "" {
		dst.ResponseType = src.ResponseType
	}
	if src.ResponseEncoding != "" {
		dst.ResponseEncoding = src.ResponseEncoding
	}
	if src.MaxRedirects != 0 {
		dst.MaxRedirects = src.MaxRedirects
	}
	if src.MaxContentLength != 0 {
		dst.MaxContentLength = src.MaxContentLength
	}
	if src.MaxBodyLength != 0 {
		dst.MaxBodyLength = src.MaxBodyLength
	}
	if src.ValidateStatus != nil {
		dst.ValidateStatus = src.ValidateStatus
	}
	if src.InterceptorOptions.RequestInterceptors != nil {
		dst.InterceptorOptions.RequestInterceptors = src.InterceptorOptions.RequestInterceptors
	}
	if src.InterceptorOptions.ResponseInterceptors != nil {
		dst.InterceptorOptions.ResponseInterceptors = src.InterceptorOptions.ResponseInterceptors
	}
	if src.OnUploadProgress != nil {
		dst.OnUploadProgress = src.OnUploadProgress
	}
	if src.OnDownloadProgress != nil {
		dst.OnDownloadProgress = src.OnDownloadProgress
	}
	if src.Proxy != nil {
		dst.Proxy = src.Proxy
	}
	if src.Cache != nil {
		dst.Cache = src.Cache
	}
	dst.Decompress = src.Decompress
}