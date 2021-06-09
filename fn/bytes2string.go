package fn

import "unsafe"

// Bytes2String 高性能转换
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
