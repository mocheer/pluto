package fn

import "encoding/base64"

// Atob 解码base64编码字符串
func Atob(s string) string {
	return string(Atob2Bytes(s))
}

// Atob2Bytes 解码base64编码字符串
func Atob2Bytes(s string) []byte {
	ret, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return ret
}
