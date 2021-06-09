package fn

import (
	"io"
	"net/http"
)

// Load 下载
func Load(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body) //body 拿到请求返回的内容
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}
