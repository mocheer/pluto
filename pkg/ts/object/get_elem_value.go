package object

import "reflect"

// GetElemValue 获取实例的反射值，如果是指针，则遍历获取到非指针实例
func GetElemValue(v any) reflect.Value {
	val := reflect.ValueOf(v)
	for val.Kind() == reflect.Ptr {
		val = val.Elem() // reflect.Indirect(v)
	}
	return val
}
