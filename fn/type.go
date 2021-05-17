package fn

import (
	"reflect"
)

// GetType 通过反射获取数据类型
// GetType(string) == "string"
// GetType(nil)
func GetType(val interface{}) reflect.Kind {
	if val == nil {
		return reflect.Invalid
	}
	return reflect.TypeOf(val).Kind()
}

// IsString
func IsString(val interface{}) bool {
	return GetType(val) == reflect.String
}
