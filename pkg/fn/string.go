package fn

import (
	"strconv"
	"strings"
)

// 类似于javascript，全局的parseInt、parseFloat
// 需要去掉首尾空格，否则可能会解析为0
// ParseFloat64
func ParseFloat64(str string) float64 {
	f64, _ := strconv.ParseFloat(strings.TrimSpace(str), 64)
	return f64
}

// ParseFloat32
func ParseFloat32(str string) float32 {
	f64, _ := strconv.ParseFloat(str, 32)
	return float32(f64)
}

// ParseInt64
func ParseInt64(str string) int64 {
	i64, _ := strconv.ParseInt(str, 10, 64)
	return i64
}

// ParseInt32 有问题，不包括浮点数解析
func ParseInt32(str string) int32 {
	i64, _ := strconv.ParseInt(str, 10, 32)
	return int32(i64)
}

// ParseInt16
func ParseInt16(str string) int16 {
	i64, _ := strconv.ParseInt(str, 10, 16)
	return int16(i64)
}

// ParseInt8
func ParseInt8(str string) int8 {
	i64, _ := strconv.ParseInt(str, 10, 8)
	return int8(i64)
}

// ParseInt 解析失败时，返回int默认值
func ParseInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

// ParseUint
// base：   数字字符串的进制，比如：2进制、10进制，如果 base 为 0，则根据字符串的前缀判断进位制（0x:16，0:8，其它:10）
// bitSize：预期数值的bit大小，用于数值上限限制，最终返回的还是int64/uint64类型
func ParseUint(str string) uint {
	u64, _ := strconv.ParseUint(str, 10, 0)
	return uint(u64)
}

// ParseHex
func ParseHex(str string) uint64 {
	u64, _ := strconv.ParseUint(str, 16, 0)
	return u64
}

// ParseBool 将字符串转换为布尔值
// 它接受真值：1, t, T, TRUE, true, True
// 它接受假值：0, f, F, FALSE, false, False.
// 其它任何值都返回一个错误
func ParseBool(str string) bool {
	b, _ := strconv.ParseBool(str)
	return b
}
