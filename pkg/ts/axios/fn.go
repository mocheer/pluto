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
