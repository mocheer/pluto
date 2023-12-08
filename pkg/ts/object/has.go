package object

import (
	"reflect"
	"strings"
)

// Has 通过判断对象是否包含某一个类型
func Has(obj any, name string) bool {
	_, ok := GetField(obj, name)
	return ok
}

// Has 通过判断对象是否包含某一个类型
func HasJsonTag(obj any, name string) bool {
	typ := getElemType(obj)
	num := typ.NumField() // 包括未导出的字段
	for i := 0; i < num; i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("json")
		if tag == name {
			return true
		}
	}
	return false
}

// GetJsonTags 通过判断对象是否包含某一个类型
func GetJsonTags(obj any) []string {
	typ := getElemType(obj)

	num := typ.NumField() // 包括未导出的字段
	fieldNames := make([]string, num)
	for i := 0; i < num; i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("json")
		fieldNames = append(fieldNames, tag)
	}
	return fieldNames
}

// GetField
func GetField(obj any, name string) (reflect.StructField, bool) {
	return getElemType(obj).FieldByName(name)
}

// GetFieldValue
func GetFieldValue(obj any, name string) reflect.Value {
	return GetElemValue(obj).FieldByName(name)
}

func GetFieldValueWithLower(obj any, name string) reflect.Value {
	return GetElemValue(obj).FieldByNameFunc(func(fileName string) bool {
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
