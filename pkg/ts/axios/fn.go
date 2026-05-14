package axios

// mergeOptions 将源请求选项合并到目标请求选项中
func mergeOptions(dst, src *AxiosOptions) {
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
	if src.Header != nil {
		dst.Header = src.Header
	}
	if src.Timeout != 0 {
		dst.Timeout = src.Timeout
	}
	if src.Auth != nil {
		dst.Auth = src.Auth
	}
	if src.MaxRedirects != 0 {
		dst.MaxRedirects = src.MaxRedirects
	}
	if src.MaxResponseContentSize != 0 {
		dst.MaxResponseContentSize = src.MaxResponseContentSize
	}
	if src.MaxRequestBodySize != 0 {
		dst.MaxRequestBodySize = src.MaxRequestBodySize
	}
	if src.ValidateStatus != nil {
		dst.ValidateStatus = src.ValidateStatus
	}
	if src.RequestInterceptors != nil {
		dst.RequestInterceptors = src.RequestInterceptors
	}
	if src.ResponseInterceptors != nil {
		dst.ResponseInterceptors = src.ResponseInterceptors
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
	dst.Uncompressed = src.Uncompressed
}
