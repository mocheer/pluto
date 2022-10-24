package fn

import (
	"reflect"
)

// New 通过反射创建结构体对象，这里最好只传递引用类型
func New(v any) any {
	typ := GetReflectType(v)
	entity := reflect.New(typ)
	return entity.Interface()
}

// NewSlice 通过反射创建结构体切片
// 这里返回的是指针,gorm的find查询需要的是指针
func NewSlice(v any) any {
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
