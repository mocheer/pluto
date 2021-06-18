package rh

import (
	"github.com/imroc/req"
)

//
func Get(url string) ([]byte, error) {
	res, err := req.New().Get(url)
	if err != nil {
		return nil, err
	}
	return res.Bytes(), nil
}

//
func Post(url string, v ...interface{}) ([]byte, error) {
	res, err := req.New().Post(url, v...)
	if err != nil {
		return nil, err
	}
	return res.Bytes(), nil
}
