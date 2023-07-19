package object

import (
	"reflect"
	"strings"
)

// GetStructFields
// 不包括忽略的字段，返回字段索引和字段名称
func GetStructFields(typ reflect.Type) ([]int, []string) {
	num := typ.NumField() // 包括未导出的字段
	fieldIndexs := []int{}
	fieldNames := []string{}
	for i := 0; i < num; i++ {
		vtyp := typ.Field(i)
		// 忽略没有导出的字段和无效值
		if !vtyp.IsExported() {
			continue
		}
		tag := vtyp.Tag.Get("json")
		if tag == "-" {
			continue
		}
		name := tag
		if name == "" {
			name = vtyp.Name
			name = strings.ToLower(name[:1]) + name[1:]
		} else {
			values := strings.Split(tag, ",")
			if len(values) > 1 { //TODO：这里还需要解析 omitempty
				name = values[0]
			}
		}
		fieldIndexs = append(fieldIndexs, i)
		fieldNames = append(fieldNames, name)
	}
	return fieldIndexs, fieldNames
}
