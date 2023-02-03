package fn

import "unsafe"

//  StringToBytes String to Bytes 高性能转换
func StringToBytes(s string) []byte {
	// x := (*[2]uintptr)(unsafe.Pointer(&s))
	// h := [3]uintptr{x[0], x[1], x[1]}
	// return *(*[]byte)(unsafe.Pointer(&h))
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
