package object

import (
	"reflect"
	"strings"
)

// GetFieldsByReflectStruct
// 返回struct所有的字段索引和字段名称， 不包括忽略的字段。
// "json:-"
func GetFieldsByReflectStruct(typ reflect.Type) ([]int, []string) {
	num := typ.NumField() // 包括未导出的字段
	fieldIndexs := []int{}
	fieldNames := []string{}
	for i := 0; i < num; i++ {
		field := typ.Field(i)
		// 忽略没有导出的字段和无效值
		if !field.IsExported() {
			continue
		}
		tag := field.Tag.Get("json")
		if tag == "-" {
			continue
		}
		name := tag
		if name == "" {
			name = field.Name
			name = strings.ToLower(name[:1]) + name[1:]
		} else {
			values := strings.Split(tag, ",")
			if len(values) > 1 { //TODO：这里还需要解析 omitempty 用于忽略空值
				name = values[0]
			}
		}
		fieldIndexs = append(fieldIndexs, i)
		fieldNames = append(fieldNames, name)
	}
	return fieldIndexs, fieldNames
}
