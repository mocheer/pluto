package rh

import (
	"github.com/imroc/req"
)

// Get 请求http数据接口
func Get(url string, v ...interface{}) ([]byte, error) {
	r := req.New()
	r.EnableInsecureTLS(true) //不校验https证书，如果校验，类似企业微信认证可能出现错误：x509: certificate signed by unknown authority
	res, err := r.Get(url, v...)
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
