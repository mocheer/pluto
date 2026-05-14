package axios

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

func createReqBody(body any) (io.Reader, int64) {
	var bodyReader io.Reader
	var bodyLength int64
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
	return bodyReader, bodyLength
}