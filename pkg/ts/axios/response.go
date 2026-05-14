package axios

import (
	"bytes"
	"encoding/json"
	"image"
	"io"
	"net/http"
	"os"
)

// Response 表示 HTTP 响应
// 默认解压缩响应体，因为响应体可能是 gzip/deflate/br/zstd 压缩的
type Response struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

// JSON 将响应体解析为 JSON 格式
func (r *Response) JSON(v interface{}) error {
	return json.Unmarshal(r.Body, v)
}

// Text 将响应体按 UTF-8 字符串返回
func (r *Response) Text() string {
	return string(r.Body)
}

// Reader 返回响应体的 io.Reader，便于流式处理
func (r *Response) Reader() io.Reader {
	return bytes.NewReader(r.Body)
}

// Save 将响应体保存到文件中
func (r *Response) Save(filename string) error {
	return os.WriteFile(filename, r.Body, 0644)
}

// Image 将响应体解析为图片数据
func (r *Response) Image() image.Image {
	img, _, err := image.Decode(bytes.NewReader(r.Body))
	if err != nil {
		panic(err)
	}
	return img
}
