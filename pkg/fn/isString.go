package fn

import "reflect"

// IsString
func IsString(v interface{}) bool {
	return GetKind(v) == reflect.String
}
