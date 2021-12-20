package fn

import (
	"reflect"
)

// GetKind 通过反射获取数据类型
// GetKind(string) == reflect.String
// GetKind(nil) == reflect.Invalid
func GetKind(v interface{}) reflect.Kind {
	return reflect.ValueOf(v).Kind()
}

// GetType(string) == "string"
// GetType(int) == "int"
// GetType(int32) == "int32"
// GetType(int64) == "int64"
func GetType(v interface{}) string {
	return GetKind(v).String()
}
