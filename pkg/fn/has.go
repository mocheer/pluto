package fn

import (
	"reflect"
	"strings"
)

// Has 通过判断对象是否包含某一个类型
func Has(obj any, name string) bool {
	_, ok := GetField(obj, name)
	return ok
}

// GetField
func GetField(obj any, name string) (reflect.StructField, bool) {
	return GetReflectType(obj).FieldByName(name)
}

func GetFieldValu(obj any, name string) reflect.Value {
	return GetReflectValue(obj).FieldByName(name)
}

func GetFieldValueWithLower(obj any, name string) reflect.Value {
	return GetReflectValue(obj).FieldByNameFunc(func(fileName string) bool {
		return strings.ToLower(fileName) == name
	})
}

// GetTag 获取结构体标签
func GetTag(obj struct{}, name, tagName string) string {
	field, ok := GetField(obj, name)
	if ok {
		return field.Tag.Get(tagName)
	}
	return ""
}
