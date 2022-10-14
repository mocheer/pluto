package fn

import "unsafe"

// B2S Byte to String 高性能转换
// 注意，当byte数据修改，string数据也会修改，容易引发问题
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
