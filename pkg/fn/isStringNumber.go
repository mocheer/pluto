package fn

import (
	"regexp"
)

// IsStringNumber 判断字符串是否是数字
func IsStringNumber(str string) bool {

	pattern := regexp.MustCompile(`^[-+]?\d*\.?\d+(?:[eE][-+]?\d+)?$`)
	return pattern.MatchString(str)
}
