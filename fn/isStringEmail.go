package fn

import (
	"regexp"
)

// IsStringEmail 判断字符串是否符合邮箱地址
func IsStringEmail(value string) bool {
	pattern := regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
	return pattern.MatchString(value)
}
