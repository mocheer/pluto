package object

import (
	"reflect"

	"github.com/samber/lo"
)

// 获取一个map的所有key值
func Keys(v any) []any {
	return lo.Map(reflect.ValueOf(v).MapKeys(), func(key reflect.Value, _ int) any {
		return key.Interface()
	})
}

// 获取一个map的所有key值
func KeysString(v any) []string {
	return lo.Map(reflect.ValueOf(v).MapKeys(), func(key reflect.Value, _ int) string {
		return key.String()
	})
}
