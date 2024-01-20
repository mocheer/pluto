package object

import (
	"reflect"
)

// GetTypeName 获取类型名称
// 可获取struct的名称,不包含包名
func GetTypeName(v any) string {
	typ := reflect.TypeOf(v)
	// 如果是指针需要指向值才能获取正确的类型
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ.Name()
}
