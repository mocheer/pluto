package axios

// Get 异步发送 HTTP GET 请求，返回 Promise
// 当通道不存在时, 请求实际为同步模式
func Get(urlStr string, options ...*AxiosOptions) *Promise {
	return Request(MethodGet, urlStr, options...)
}

// Post 异步发送 HTTP POST 请求，返回 Promise
func Post(urlStr string, options ...*AxiosOptions) *Promise {
	return Request(MethodPost, urlStr, options...)
}

// Put 异步发送 HTTP PUT 请求，返回 Promise
func Put(urlStr string, options ...*AxiosOptions) *Promise {
	return Request(MethodPut, urlStr, options...)
}

// Delete 异步发送 HTTP DELETE 请求，返回 Promise
func Delete(urlStr string, options ...*AxiosOptions) *Promise {
	return Request(MethodDelete, urlStr, options...)
}

// Head 异步发送 HTTP HEAD 请求，返回 Promise
func Head(urlStr string, options ...*AxiosOptions) *Promise {
	return Request(MethodHead, urlStr, options...)
}

// Options 异步发送 HTTP OPTIONS 请求，返回 Promise
func Options(urlStr string, options ...*AxiosOptions) *Promise {
	return Request(MethodOptions, urlStr, options...)
}

// Patch 异步发送 HTTP PATCH 请求，返回 Promise
func Patch(urlStr string, options ...*AxiosOptions) *Promise {
	return Request(MethodPatch, urlStr, options...)
}

// Get 发送 HTTP GET 请求
func GetSync(urlStr string, options ...*AxiosOptions) (*Response, error) {
	return RequestSync(MethodGet, urlStr, options...)
}

// PostSync 发送 HTTP POST 请求，包含请求体
func PostSync(urlStr string, options ...*AxiosOptions) (*Response, error) {
	return RequestSync(MethodPost, urlStr, options...)
}

// Put 发送 HTTP PUT 请求
func PutSync(urlStr string, options ...*AxiosOptions) (*Response, error) {
	return RequestSync(MethodPut, urlStr, options...)
}

// Delete 发送 HTTP DELETE 请求
func DeleteSync(urlStr string, options ...*AxiosOptions) (*Response, error) {
	return RequestSync(MethodDelete, urlStr, options...)
}

// Head 发送 HTTP HEAD 请求
func HeadSync(urlStr string, options ...*AxiosOptions) (*Response, error) {
	return RequestSync(MethodHead, urlStr, options...)
}

// Options 发送 HTTP OPTIONS 请求
func OptionsSync(urlStr string, options ...*AxiosOptions) (*Response, error) {
	return RequestSync(MethodOptions, urlStr, options...)
}

// Patch 发送 HTTP PATCH 请求
func PatchSync(urlStr string, options ...*AxiosOptions) (*Response, error) {
	return RequestSync(MethodPatch, urlStr, options...)
}

// RequestSync 发送 HTTP 请求，使用默认客户端
func RequestSync(method MethodType, urlStr string, configs ...*AxiosOptions) (*Response, error) {
	axiosOptions := &AxiosOptions{
		Method: method,
		URL:    urlStr,
	}
	for _, opt := range configs {
		mergeOptions(axiosOptions, opt)
	}
	return defaultClient.Request(axiosOptions)
}

// Request 异步发送 HTTP 请求，返回 Promise
func Request(method MethodType, urlStr string, configs ...*AxiosOptions) *Promise {
	axiosOptions := &AxiosOptions{
		Method: method,
		URL:    urlStr,
	}
	for _, opt := range configs {
		mergeOptions(axiosOptions, opt)
	}
	return defaultClient.RequestAsync(axiosOptions)
}
