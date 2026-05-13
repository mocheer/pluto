package axios

// Get 异步发送 HTTP GET 请求，返回 Promise
func Get(urlStr string, options ...*RequestOptions) *Promise {
	return Request(MethodGet, urlStr, options...)
}

// Post 异步发送 HTTP POST 请求，返回 Promise
func Post(urlStr string, body interface{}, options ...*RequestOptions) *Promise {
	mergedOptions := mergeBodyIntoOptions(body, options)
	return Request(MethodPost, urlStr, mergedOptions)
}

// Put 异步发送 HTTP PUT 请求，返回 Promise
func Put(urlStr string, body interface{}, options ...*RequestOptions) *Promise {
	mergedOptions := mergeBodyIntoOptions(body, options)
	return Request(MethodPut, urlStr, mergedOptions)
}

// Delete 异步发送 HTTP DELETE 请求，返回 Promise
func Delete(urlStr string, options ...*RequestOptions) *Promise {
	return Request(MethodDelete, urlStr, options...)
}

// Head 异步发送 HTTP HEAD 请求，返回 Promise
func Head(urlStr string, options ...*RequestOptions) *Promise {
	return Request(MethodHead, urlStr, options...)
}

// Options 异步发送 HTTP OPTIONS 请求，返回 Promise
func Options(urlStr string, options ...*RequestOptions) *Promise {
	return Request(MethodOptions, urlStr, options...)
}

// Patch 异步发送 HTTP PATCH 请求，返回 Promise
func Patch(urlStr string, body interface{}, options ...*RequestOptions) *Promise {
	mergedOptions := mergeBodyIntoOptions(body, options)
	return Request(MethodPatch, urlStr, mergedOptions)
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

// Get 发送 HTTP GET 请求
func GetSync(urlStr string, options ...*RequestOptions) (*Response, error) {
	return RequestSync(MethodGet, urlStr, options...)
}

// PostSync 发送 HTTP POST 请求，包含请求体
func PostSync(urlStr string, body interface{}, options ...*RequestOptions) (*Response, error) {
	mergedOptions := mergeBodyIntoOptions(body, options)
	return RequestSync(MethodPost, urlStr, mergedOptions)
}

// Put 发送 HTTP PUT 请求
func PutSync(urlStr string, body interface{}, options ...*RequestOptions) (*Response, error) {
	mergedOptions := mergeBodyIntoOptions(body, options)
	return RequestSync(MethodPut, urlStr, mergedOptions)
}

// Delete 发送 HTTP DELETE 请求
func DeleteSync(urlStr string, options ...*RequestOptions) (*Response, error) {
	return RequestSync(MethodDelete, urlStr, options...)
}

// Head 发送 HTTP HEAD 请求
func HeadSync(urlStr string, options ...*RequestOptions) (*Response, error) {
	return RequestSync(MethodHead, urlStr, options...)
}

// Options 发送 HTTP OPTIONS 请求
func OptionsSync(urlStr string, options ...*RequestOptions) (*Response, error) {
	return RequestSync(MethodOptions, urlStr, options...)
}

// Patch 发送 HTTP PATCH 请求
func PatchSync(urlStr string, body interface{}, options ...*RequestOptions) (*Response, error) {
	mergedOptions := mergeBodyIntoOptions(body, options)
	return RequestSync("PATCH", urlStr, mergedOptions)
}

// RequestSync 发送 HTTP 请求，使用默认客户端
func RequestSync(method MethodType, urlStr string, options ...*RequestOptions) (*Response, error) {
	reqOptions := &RequestOptions{
		Method:           MethodGet,
		URL:              urlStr,
		Timeout:          1000,
		ResponseType:     "json",
		ResponseEncoding: "utf8",
		MaxContentLength: 2 * 1024 * 1024,
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

// Request 异步发送 HTTP 请求，返回 Promise
func Request(method MethodType, urlStr string, options ...*RequestOptions) *Promise {
	reqOptions := &RequestOptions{
		Method:           MethodGet,
		URL:              urlStr,
		Timeout:          1000,
		ResponseType:     "json",
		ResponseEncoding: "utf8",
		MaxContentLength: 2 * 1024 * 1024,
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

	return defaultClient.RequestAsync(reqOptions)
}

// SetBaseURL 设置默认客户端的 BaseURL
func SetBaseURL(baseURL string) {
	defaultClient.BaseURL = baseURL
}
