package axios

// SetBaseURL 设置默认客户端的 BaseURL
func SetBaseURL(baseURL string) {
	defaultClient.Options.BaseURL = baseURL
}

// SetLogger 设置默认客户端的日志记录器
func SetLogger(logger Logger) {
	defaultClient.logger = logger
}

// SetMaxPendingRequests 设置默认客户端的最大并发请求数
func SetMaxPendingRequests(size int) {
	defaultClient.SetMaxPendingRequests(size)
}
