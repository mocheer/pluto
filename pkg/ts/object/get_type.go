package object

// GetType
// 没办法判断time.Time
// GetType(string) == "string"
// GetType(int) == "int"
// GetType(int32) == "int32"
// GetType(int64) == "int64"
func GetType(v any) string {
	return GetKind(v).String()
}
