package js

import (
	"encoding/base64"

	"github.com/mocheer/pluto/fn"
)

// Btoa 解码base64字符串
func Btoa(s string) string {
	return base64.StdEncoding.EncodeToString(fn.S2B(s))
}
