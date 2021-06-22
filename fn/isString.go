package fn

import "reflect"

// IsPtr 是否为指针
func IsPtr(v interface{}) bool {
	return GetKind(v) == reflect.Ptr
}
