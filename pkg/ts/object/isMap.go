package object

import "reflect"

// IsInt
func IsMap(v any) bool {
	return GetKind(v) == reflect.Map
}
