package fn

import "reflect"

// IsInt
func IsMap(v interface{}) bool {
	return GetKind(v) == reflect.Map
}
