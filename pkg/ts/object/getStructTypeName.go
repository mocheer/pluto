package object

import (
	"reflect"
)

// 获取一个struct的类型名称
func GetStructTypeName(v any) string {
	return reflect.TypeOf(v).Name()
}
