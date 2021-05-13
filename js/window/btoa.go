package window

import "encoding/base64"

func Btoa(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
