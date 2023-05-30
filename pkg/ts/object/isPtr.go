package object

import "reflect"

// IsPtr 是否为指针
func IsPtr(v any) bool {
	return GetKind(v) == reflect.Ptr
}
