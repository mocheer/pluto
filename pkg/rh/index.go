package rh

import (
	"github.com/imroc/req"
)

// Get 请求http数据接口
func Get(url string) ([]byte, error) {
	res, err := req.New().Get(url)
	if err != nil {
		return nil, err
	}
	return res.Bytes(), nil
}

// Post 请求http数据接口
func Post(url string, v ...interface{}) ([]byte, error) {
	res, err := req.New().Post(url, v...)
	if err != nil {
		return nil, err
	}
	return res.Bytes(), nil
}
