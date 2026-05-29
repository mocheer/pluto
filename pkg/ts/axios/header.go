package axios

import "encoding/base64"

type Header map[string]string

// SetAuth 设置认证头
func (h Header) SetAuth(auth *Auth) {
	if auth == nil {
		return
	}
	h["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString([]byte(auth.Username+":"+auth.Password))
}

// RandomUserAgent
func (h Header) RandomUserAgent() {
	h["User-Agent"] = RandomUserAgent()
}
