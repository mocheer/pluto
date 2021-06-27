package jsg

import "encoding/base64"

// Atob base64编码字符串
func Atob(s string) string {
	ret, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return string(ret)
}
