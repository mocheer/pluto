package fn

const (
	String = iota
	Int32
	Int64
	Unknown
)

func GetTypeEnum(value any) int {
	switch value.(type) {
	case string:
		return String
	case int32:
		return Int32
	case int64:
		return Int64
	default:
		return Unknown
	}
}
