package fn

import "reflect"

// IsInt
func IsInt(v any) bool {
	return GetKind(v) == reflect.Int
}
