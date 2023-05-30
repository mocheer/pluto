package object

import "reflect"

// Each 遍历任意类型的切片
func Each(arr any, callback func(value any)) {
	s := reflect.ValueOf(arr)
	for i := 0; i < s.Len(); i++ {
		callback(s.Index(i).Interface())
	}
}
