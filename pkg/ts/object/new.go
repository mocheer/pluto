package object

import (
	"reflect"
)

// New
// struct类型: 通过反射创建结构体对象,返回的是指向结构体的指针
// pointer类型：通过反射获取指针实际的类型，返回指向对应类型的指针
func New(v any) any {
	typ := getElemType(v)
	return reflect.New(typ).Interface()
}

// NewFunc
func NewFunc(v any) func() any {
	typ := getElemType(v)
	return func() any {
		return reflect.New(typ).Interface()
	}
}

// NewSlice 通过反射创建结构体切片的指针
// struct类型: 通过反射创建结构体切片对象,返回的是指向切片的指针
func NewSlice(v any) any {
	typ := getElemType(v)
	entity := reflect.New(reflect.SliceOf(typ))
	return entity.Interface()
}

// getElemType 获取数据类型(不包括指针)
func getElemType(v any) reflect.Type {
	val := reflect.TypeOf(v)
	for val.Kind() == reflect.Ptr {
		val = val.Elem() // reflect.Indirect(v)
	}
	return val
}
