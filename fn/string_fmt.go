package fn

import (
	"fmt"
	"regexp"
)

// FmtString
// @param str
// @param data
// @example FmtString("{a}",map[string]interfacle{}{"a":1})
func FmtString(src string, data map[string]interface{}) string {
	// 匹配花括号内的字符串`{xxx}`，常用于字符串格式化替换
	return regexp.MustCompile(`{([^}]+)}`).ReplaceAllStringFunc(src, func(key string) string {
		// 这里的key包含括号
		val := data[key[1:len(key)-1]]
		if val == nil {
			return ""
		}
		return fmt.Sprint(val)
	})
}

// // StringFormat
// // @param str
// // @param data
// // @example ([]byte(`{a} b cd{e} fg {h}`),map[string]interface{}{ "a":"1","b":2,"e":5})
// func StringFormat(str []byte, data map[string]interface{}) []byte {
// 	l := len(str)
// 	at := -1
// 	var c byte
// 	var buffer bytes.Buffer
// read_next:
// 	at++
// 	if at >= l {
// 		return buffer.Bytes()
// 	}
// 	c = str[at]
// 	if c == '{' {
// 		goto read_value
// 	}
// 	if c == '\\' {
// 		goto read_next
// 	}
// 	buffer.WriteByte(c)
// 	goto read_next
// read_value:
// 	var key []byte
// 	for {
// 		at++
// 		if at >= l {
// 			buffer.Write(key)
// 			return buffer.Bytes()
// 		}
// 		c = str[at]
// 		if c == '}' {
// 			break
// 		}
// 		if c == '\\' {
// 			continue
// 		}
// 		key = append(key, c)
// 	}
// 	value := data[string(key)]
// 	if value != nil {
// 		val := fmt.Sprint(value)
// 		buffer.WriteString(val)
// 	}
// 	goto read_next
// }
