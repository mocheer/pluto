package fn

import "unsafe"

// BytesToString Byte to String 高性能转换
// 注意，当byte数据修改，string数据也会修改，容易引发问题
func BytesToString(b []byte) string {
	// return *(*string)(unsafe.Pointer(&b))
	return unsafe.String(&b[0], len(b))
}
