package fn

import (
	"bytes"
	"fmt"
)

// FormatBytes 将 []byte 切片格式化为类似 []byte{v1, v2, v3, ...} 的字符串
func FormatBytes(data []byte) string {
	var buf bytes.Buffer
	buf.WriteString("[]byte{")
	for i, b := range data {
		if i > 0 {
			buf.WriteString(", ")
		}
		// 使用 %d 来格式化字节为十进制整数
		buf.WriteString(fmt.Sprintf("%d", b))
	}
	buf.WriteString("}")
	return buf.String()
}
