package js

import "encoding/base64"

func Atob(s string) string {
	ret, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return string(ret)
}
