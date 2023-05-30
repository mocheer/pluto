package object

import "reflect"

// IsInt
func IsInt(v any) bool {
	return GetKind(v) == reflect.Int
}

// IsInt64
func IsInt64(v any) bool {
	return GetKind(v) == reflect.Int64
}

// IsInt32
func IsInt32(v any) bool {
	return GetKind(v) == reflect.Int32
}

// IsInt8
func IsInt8(v any) bool {
	return GetKind(v) == reflect.Int8
}
