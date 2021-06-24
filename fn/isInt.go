package fn

import "reflect"

// IsInt
func IsInt(v interface{}) bool {
	return GetKind(v) == reflect.Int
}
