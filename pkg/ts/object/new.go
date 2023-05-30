package object

import (
	"reflect"
)

// NewStructPointer 通过反射创建结构体对象，这里最好只传递引用类型
func NewStructPointer(v any) any {
	typ := GetReflectType(v)
	entity := reflect.New(typ)
	return entity.Interface()
}

// NewSlicePointer 通过反射创建结构体切片的指针
// gorm的find查询需要的是指针
func NewSlicePointer(v any) any {
	typ := GetReflectType(v)
	entity := reflect.New(reflect.SliceOf(typ))
	return entity.Interface()
}

// GetType 获取数据类型(不包括指针)
func GetReflectType(v any) reflect.Type {
	val := reflect.TypeOf(v)
	for val.Kind() == reflect.Ptr {
		val = val.Elem() // reflect.Indirect(v)
	}
	return val
}

// GetType 获取数据值(不包括指针)
func GetReflectValue(v any) reflect.Value {
	val := reflect.ValueOf(v)
	for val.Kind() == reflect.Ptr {
		val = val.Elem() // reflect.Indirect(v)
	}
	return val
}
