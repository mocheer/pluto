package fn

import (
	"encoding/base64"
)

// Btoa 编码base64字符串
func Btoa(s string) string {
	return BtoaBytes(S2B(s))
}

//
func BtoaBytes(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}
