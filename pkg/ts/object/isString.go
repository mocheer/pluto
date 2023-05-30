package object

import "reflect"

// IsString
func IsString(v any) bool {
	return GetKind(v) == reflect.String
}
