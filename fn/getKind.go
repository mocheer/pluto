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
