package fn

import (
	"reflect"
)

// GetKind 通过反射获取数据类型
// GetKind(string) == reflect.String
// GetKind(nil) == reflect.Invalid
func GetKind(v any) reflect.Kind {
	return reflect.ValueOf(v).Kind()
}
