package fn

// IsStringBlank 是否空字符串
func IsStringBlank(str string) bool {
	for _, r := range str {
		switch r {
		case 9:
		case 10:
		case 13:
		case 32:
		default:
			return false
		}
	}
	return true
}
