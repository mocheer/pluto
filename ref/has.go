package ref

import "reflect"

// Has 判断对象是否包含某一个类型
func Has(obj interface{}, key string) bool {
	t := reflect.TypeOf(obj).Elem()
	_, flag := t.FieldByName(key)
	return flag
}
