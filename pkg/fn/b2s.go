package fn

import "unsafe"

// B2S Byte to String 高性能转换
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
