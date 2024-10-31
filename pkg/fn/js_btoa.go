package fn

import (
	"encoding/base64"
	"fmt"
)

// Btoa 编码base64字符串
func Btoa(s string) string {
	return BtoaBytes(StringToBytes(s))
}

func BtoaBytes(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

// BtoaWithURI
// typeName eg. "image/jpeg"
// "data:image/jpeg;base64,%s"
func BtoaWithURI(typeName string, bytes []byte) string {
	return fmt.Sprintf("data:%s;base64,%s", typeName, BtoaBytes(bytes))
}
