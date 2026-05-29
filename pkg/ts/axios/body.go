package axios

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

// createReqBody 创建请求体
func createReqBody(body any) (bodyReader io.Reader, bodyLength int64) {
	switch v := body.(type) {
	case string:
		bodyReader = strings.NewReader(v)
		bodyLength = int64(len(v))
	case []byte:
		bodyReader = bytes.NewReader(v)
		bodyLength = int64(len(v))
	default:
		jsonBody, err := json.Marshal(v)
		if err != nil {
			return nil, 0
		}
		bodyReader = bytes.NewBuffer(jsonBody)
		bodyLength = int64(len(jsonBody))
	}
	return
}
