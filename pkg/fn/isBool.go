package fn

import "reflect"

// IsBool
func IsBool(v interface{}) bool {
	return GetKind(v) == reflect.Bool
}
