package object

import (
	"reflect"
)

// GetTypeName 获取类型名称
// 可获取struct的名称
func GetTypeName(v any) string {
	return reflect.TypeOf(v).Name()
}
