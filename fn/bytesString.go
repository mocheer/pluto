package fn

import "unsafe"

//  String2Bytes 高性能转换
func String2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// Bytes2String 高性能转换
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
