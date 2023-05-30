package object

import "reflect"

// IsTime 是否为time.Time类型
func IsTime(v any) bool {
	return reflect.TypeOf(v).String() == "time.Time"
}
