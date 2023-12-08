package object

import (
	"reflect"
	"strings"
)

// FindFieldValue
func FindFieldValue(obj any, name string, tagName string) reflect.Value {
	structValue := GetElemValue(obj)
	fieldIndex := FindFieldIndex(structValue.Type(), name, tagName)
	fieldValue := structValue.Field(fieldIndex)
	return fieldValue
}

// FindFieldIndex
// tagName一般为json
// typ不能是指针类型，必须是结构体类型
func FindFieldIndex(typ reflect.Type, tagVal string, tagName string) int {
	num := typ.NumField() // 包括未导出的字段
	for i := 0; i < num; i++ {
		field := typ.Field(i)
		// 忽略没有导出的字段和无效值
		if !field.IsExported() {
			continue
		}
		jsonTagValue := field.Tag.Get(tagName)
		if jsonTagValue == "-" {
			continue
		}
		name := jsonTagValue
		if name == "" {
			name = field.Name
			name = strings.ToLower(name[:1]) + name[1:]
		} else {
			values := strings.Split(jsonTagValue, ",")
			if len(values) > 1 {
				name = values[0]
			}
		}
		if name == tagVal {
			return i
		}
	}
	return -1
}
