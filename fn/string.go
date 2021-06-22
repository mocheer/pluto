package fn

import (
	"strconv"
)

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
