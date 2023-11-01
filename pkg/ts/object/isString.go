package object

import "reflect"

// IsString
// 即使是继承string，也会为true
func IsString(v any) bool {
	return GetKind(v) == reflect.String
}
