package fn

import (
	"reflect"
)

// GetKind 通过反射获取数据类型
// GetKind(string) == "string"
// GetKind(nil)
func GetKind(v interface{}) reflect.Kind {
	if v == nil {
		return reflect.Invalid
	}
	return reflect.TypeOf(v).Kind()
}

// IsString
func IsString(v interface{}) bool {
	return GetKind(v) == reflect.String
}

// IsPtr 是否为指针
func IsPtr(v interface{}) bool {
	return GetKind(v) == reflect.Ptr
}

// 获取一个struct的类型名称
func GetStructTypeName(v interface{}) string {
	return reflect.TypeOf(v).Name()
}
