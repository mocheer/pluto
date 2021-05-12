package img

import "encoding/base64"

// Base64ToBytes 将base64转成bytes
func Base64ToBytes(data string) ([]byte, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(data)
	return decodeBytes, err
}

// BytesToBase64 将bytes转成base64
func BytesToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
