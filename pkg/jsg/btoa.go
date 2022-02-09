package jsg

import (
	"encoding/base64"

	"github.com/mocheer/pluto/pkg/fn"
)

// Btoa 编码base64字符串
func Btoa(s string) string {
	return BtoaBytes(fn.S2B(s))
}

//
func BtoaBytes(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}
