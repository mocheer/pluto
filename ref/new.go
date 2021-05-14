package ref

import "reflect"

// New 通过反射创建结构体对象
// @param v 目前支持 结构体指针对象
func New(v interface{}) interface{} {
	val := reflect.ValueOf(v)
	typ := reflect.Indirect(val).Type()
	n := reflect.New(typ)
	return n.Interface()
}
