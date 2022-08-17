package fn

import "reflect"

// IsTime 是否为time.Time类型
func IsTime(v any) bool {
	return GetKind(v) == reflect.Array
}
