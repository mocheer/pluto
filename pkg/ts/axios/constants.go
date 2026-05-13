package axios

import "net/http"

// defaultClient 是默认的 HTTP 客户端实例
var defaultClient = &Client{HTTPClient: &http.Client{}, Logger: NewLogger(LevelDebug), queueChan: make(chan struct{}, 100)}

// MethodType 表示 HTTP 请求方法类型
type MethodType string

const (
	// MethodGet 表示 HTTP GET 请求方法
	MethodGet MethodType = "GET"
	// MethodPost 表示 HTTP POST 请求方法
	MethodPost MethodType = "POST"
	// MethodPut 表示 HTTP PUT 请求方法
	MethodPut MethodType = "PUT"
	// MethodDelete 表示 HTTP DELETE 请求方法
	MethodDelete MethodType = "DELETE"
	// MethodPatch 表示 HTTP PATCH 请求方法
	MethodPatch MethodType = "PATCH"
	// MethodHead 表示 HTTP HEAD 请求方法
	MethodHead MethodType = "HEAD"
	// MethodOption 表示 HTTP OPTIONS 请求方法
	MethodOptions MethodType = "OPTIONS"
)
