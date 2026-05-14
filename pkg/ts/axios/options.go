package axios

import "time"

// AxiosOptions 包含 HTTP 请求的所有配置选项
type AxiosOptions struct {
	Method             MethodType                        // 请求方法
	URL                string                            // 请求 URL，包含查询参数
	BaseURL            string                            // 基础 URL，用于构建完整 URL
	Params             map[string]string                 // 查询参数
	Body               interface{}                       // 请求体数据，可以是[]byte、结构体、map、slice等类型，TODO 实现BodyReader类型
	Header             Header                            // 请求头
	Timeout            time.Duration                     // 请求超时时间，默认 1000ms，单位秒
	Auth               *Auth                             // 基本认证信息，包含用户名和密码
	ResponseType       string                            // 响应体类型，默认 JSON 格式
	ResponseEncoding   string                            // 响应体编码，默认 UTF-8
	MaxRedirects       int                               // 最大重定向次数，默认 21 次
	MaxContentLength   int64                             // 最大响应内容长度，默认 0 表示不限制
	MaxBodyLength      int64                             // 最大请求体长度，默认 0 表示不限制
	Decompress         bool                              // 是否解压缩响应体，默认 false
	ValidateStatus     func(int) bool                    // 自定义状态码验证函数，默认 nil
	InterceptorOptions InterceptorOptions                // 请求和响应拦截器配置
	Proxy              *Proxy                            // HTTP 代理配置
	OnUploadProgress   func(bytesRead, totalBytes int64) // 上传进度回调, 实际上是数据读取进度，并非网络传输进度
	OnDownloadProgress func(bytesRead, totalBytes int64) // 下载进度回调，实际上是数据读取进度，并非网络传输进度
	Cache              *RequestCacheOptions              // 请求缓存配置
}

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
