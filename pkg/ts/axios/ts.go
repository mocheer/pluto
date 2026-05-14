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
