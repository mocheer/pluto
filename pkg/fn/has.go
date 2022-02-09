package fn

import "reflect"

// Has 通过判断对象是否包含某一个类型
func Has(obj interface{}, name string) bool {
	_, ok := GetField(obj, name)
	return ok
}

// GetField
func GetField(obj interface{}, name string) (reflect.StructField, bool) {
	return GetReflectType(obj).FieldByName(name)
}

// GetTag 获取结构体标签
func GetTag(obj struct{}, name, tagName string) string {
	field, ok := GetField(obj, name)
	if ok {
		return field.Tag.Get(tagName)
	}
	return ""
}
