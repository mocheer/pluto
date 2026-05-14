package axios

import (
	"encoding/json"
	"io"
	"net/http"
)

// Response 表示 HTTP 响应
type Response struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

// RequestInterceptors 是请求拦截器函数列表
type RequestInterceptors []func(*http.Request) error

// ResponseInterceptors 是响应拦截器函数列表
type ResponseInterceptors []func(*http.Response) error

// InterceptorOptions 包含请求和响应拦截器配置
type InterceptorOptions struct {
	RequestInterceptors  RequestInterceptors
	ResponseInterceptors ResponseInterceptors
}

// AxiosOptions 包含 HTTP 请求的所有配置选项
type AxiosOptions struct {
	Method             MethodType                        // 请求方法
	URL                string                            // 请求 URL，包含查询参数
	BaseURL            string                            // 基础 URL，用于构建完整 URL
	Params             map[string]string                 // 查询参数
	Body               interface{}                       // 请求体数据，可以是[]byte、结构体、map、slice等类型，TODO 实现BodyReader类型
	Headers            map[string]string                 // 请求头
	Timeout            int                               // 请求超时时间，单位秒
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

// Proxy 表示 HTTP 代理配置
type Proxy struct {
	Protocol string
	Host     string
	Port     int
	Auth     *Auth
}

// Auth 表示 HTTP 基本认证信息
type Auth struct {
	Username string
	Password string
}

// ProgressReader 是支持上传进度回调的读取器
type ProgressReader struct {
	reader     io.Reader
	total      int64
	read       int64
	onProgress func(bytesRead, totalBytes int64)
}

// ProgressWriter 是支持下载进度回调的写入器
type ProgressWriter struct {
	writer     io.Writer
	total      int64
	written    int64
	onProgress func(bytesWritten, totalBytes int64)
}

// Read 从 ProgressReader 中读取数据并触发进度回调
func (pr *ProgressReader) Read(p []byte) (int, error) {
	n, err := pr.reader.Read(p)
	pr.read += int64(n)
	if pr.onProgress != nil {
		pr.onProgress(pr.read, pr.total)
	}
	return n, err
}

// Write 向 ProgressWriter 中写入数据并触发进度回调
func (pw *ProgressWriter) Write(p []byte) (int, error) {
	n, err := pw.writer.Write(p)
	pw.written += int64(n)
	if pw.onProgress != nil {
		pw.onProgress(pw.written, pw.total)
	}
	return n, err
}

// JSON 将响应体解析为 JSON 格式
func (r *Response) JSON(v interface{}) error {
	return json.Unmarshal(r.Body, v)
}
