package fn

import (
	"reflect"
)

// New 通过反射创建结构体对象
func New(v interface{}) interface{} {
	typ := GetReflectType(v)
	entity := reflect.New(typ)
	return entity.Interface()
}

// NewSlice 通过反射创建结构体对象
func NewSlice(v interface{}) interface{} {
	typ := GetReflectType(v)
	entity := reflect.New(reflect.SliceOf(typ))
	return entity.Interface()
}

// GetType 获取类型
func GetReflectType(v interface{}) reflect.Type {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem() // reflect.Indirect(v)
	}
	return val.Type()
}
