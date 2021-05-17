package fn

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/mocheer/pluto/reg"
)

// FmtString
// @param str
// @param data
// @example FmtString("{a}",map[string]interfacle{}{"a":1})
func FmtString(src string, data map[string]interface{}) string {
	return regexp.MustCompile(reg.Brace).ReplaceAllStringFunc(src, func(key string) string {
		// 这里的key包含括号
		val := data[key[1:len(key)-1]]
		if val == nil {
			return ""
		}
		return fmt.Sprint(val)
	})
}

// ParseFloat64
func ParseFloat64(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

// ParseFloat32
func ParseFloat32(str string) (float32, error) {
	f64, err := strconv.ParseFloat(str, 32)
	return float32(f64), err
}

// ParseInt64
func ParseInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// ParseInt32
func ParseInt32(str string) (int32, error) {
	i64, err := strconv.ParseInt(str, 10, 32)
	return int32(i64), err
}

// ParseInt16
func ParseInt16(str string) (int16, error) {
	i64, err := strconv.ParseInt(str, 10, 16)
	return int16(i64), err
}

// ParseInt8
func ParseInt8(str string) (int8, error) {
	i64, err := strconv.ParseInt(str, 10, 8)
	return int8(i64), err
}

// ParseInt
func ParseInt(str string) (int, error) {
	return strconv.Atoi(str)
}

// ParseUint64
func ParseUint64(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 0)
}

// ParseUint
func ParseUint(str string) (uint, error) {
	u64, err := strconv.ParseUint(str, 10, 0)
	return uint(u64), err
}

// ParseHex
func ParseHex(str string) (uint64, error) {
	return strconv.ParseUint(str, 0, 0) //如果 base 为 0，则根据字符串的前缀判断进位制（0x:16，0:8，其它:10）
}

// Bool 将字符串转换为布尔值
// 它接受真值：1, t, T, TRUE, true, True
// 它接受假值：0, f, F, FALSE, false, False.
// 其它任何值都返回一个错误
func ParseBool(str string) (bool, error) {
	return strconv.ParseBool(str)
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
