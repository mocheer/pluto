package fetch

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// Response 表示 HTTP 响应，类似 fetch 的 Response 对象。
type Response struct {
	StatusCode int
	Status     string // e.g. "200 OK"
	Headers    http.Header
	Body       []byte
	RequestURL string
	err        error // 内部错误，供 Fetch 返回
}

// Err 返回处理过程中出现的网络错误（如超时、DNS错误），
// 注意：HTTP 错误状态码（4xx、5xx）不视为 Go error，而是正常的响应。
func (r *Response) Err() error {
	return r.err
}

// Text 将响应体按 UTF-8 字符串返回
func (r *Response) Text() string {
	return string(r.Body)
}

// JSON 将响应体解析为 JSON，存入 v
func (r *Response) JSON(v interface{}) error {
	return json.Unmarshal(r.Body, v)
}

// Reader 返回响应体的 io.Reader，便于流式处理
func (r *Response) Reader() io.Reader {
	return strings.NewReader(string(r.Body))
}
