package fn

import "reflect"

// IsString
func IsString(v any) bool {
	return GetKind(v) == reflect.String
}
