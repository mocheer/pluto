package fn

import "unsafe"

//  StringBytes 高性能转换
func StringBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// BytesString 高性能转换
func BytesString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
