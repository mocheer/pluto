package fn

import "reflect"

// ToSliceReflectValue
func ToSliceReflectValue(args []interface{}) []reflect.Value {
	in := make([]reflect.Value, len(args))
	for k, param := range args {
		in[k] = reflect.ValueOf(param)
	}
	return in
}

// Apply 通过反射支持字符串直接调用函数
func Apply(obj_func interface{}, args []interface{}) []reflect.Value {
	return reflect.ValueOf(obj_func).Call(ToSliceReflectValue(args))
}

// Call 通过反射支持字符串直接调用函数
func Call(obj_func interface{}, args ...interface{}) []reflect.Value {
	return reflect.ValueOf(obj_func).Call(ToSliceReflectValue(args))
}

// CallMethod
// @param obj 为结构体指针
func CallMethod(obj interface{}, methodName string, args ...interface{}) []reflect.Value {
	return reflect.ValueOf(obj).Elem().MethodByName(methodName).Call(ToSliceReflectValue(args))
}
