package fn

import (
	"encoding/base64"
	"strings"
)

// Atob 解码base64编码字符串
func Atob(s string) string {
	return string(Atob2Bytes(s))
}

// Atob2Bytes 解码base64编码字符串
// 如果是`data:image/jpeg;base64,{s}` 需要去掉前缀
func Atob2Bytes(s string) []byte {
	ret, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return ret
}

// Atob2BytesWithURI
func Atob2BytesWithURI(s string) []byte {
	if strings.Contains(s, "data:image") {
		s = s[strings.Index(s, ",")+1:]
	}
	return Atob2Bytes(s)
}
