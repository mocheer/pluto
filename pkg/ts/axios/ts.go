package axios

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"
)

// Response 表示 HTTP 响应
type Response struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

// Promise 表示异步 HTTP 请求的 Promise
type Promise struct {
	response *Response
	err      error
	then     func(*Response)
	catch    func(error)
	finally  func()
	done     chan struct{}
	mu       sync.Mutex
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

// RequestOptions 包含 HTTP 请求的所有配置选项
type RequestOptions struct {
	Method             MethodType
	URL                string
	BaseURL            string
	Params             map[string]string
	Body               interface{}
	Headers            map[string]string
	Timeout            int
	Auth               *Auth
	ResponseType       string
	ResponseEncoding   string
	MaxRedirects       int
	MaxContentLength   int64
	MaxBodyLength      int64
	Decompress         bool
	ValidateStatus     func(int) bool
	InterceptorOptions InterceptorOptions
	Proxy              *Proxy
	OnUploadProgress   func(bytesRead, totalBytes int64)
	OnDownloadProgress func(bytesRead, totalBytes int64)
	LogLevel           LogLevel
	Cache              *RequestCacheOptions
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

// Then 注册 Promise 成功时的回调函数
func (p *Promise) Then(fn func(*Response)) *Promise {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.response != nil && p.err == nil {
		fn(p.response)
	} else {
		p.then = fn
	}
	return p
}

// Catch 注册 Promise 失败时的回调函数
func (p *Promise) Catch(fn func(error)) *Promise {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.err != nil {
		fn(p.err)
	} else {
		p.catch = fn
	}
	return p
}

// Finally 注册 Promise 完成时的回调函数（无论成功或失败）
func (p *Promise) Finally(fn func()) {
	p.mu.Lock()

	if p.response != nil || p.err != nil {
		p.mu.Unlock()
		fn()
	} else {
		p.finally = fn
		p.mu.Unlock()
	}

	<-p.done
}

// NewPromise 创建一个新的 Promise
func NewPromise() *Promise {
	return &Promise{
		done: make(chan struct{}),
	}
}

// resolve 解析 Promise，设置响应或错误并触发回调
func (p *Promise) resolve(resp *Response, err error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.response = resp
	p.err = err

	if p.then != nil && err == nil {
		p.then(resp)
	}
	if p.catch != nil && err != nil {
		p.catch(err)
	}
	if p.finally != nil {
		p.finally()
	}

	close(p.done)
}
