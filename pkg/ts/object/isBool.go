package object

import "reflect"

// IsBool
func IsBool(v any) bool {
	return GetKind(v) == reflect.Bool
}
