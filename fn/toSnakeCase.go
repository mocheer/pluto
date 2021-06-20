package fn

import (
	"regexp"
	"strings"
)

// ToSnakeCase
func ToSnakeCase(val string) string {
	// 匹配驼峰转成短连接符
	return regexp.MustCompile(`([a-z])([A-Z])`).ReplaceAllStringFunc(val, func(s string) string {
		return s[:1] + "-" + strings.ToLower(s[1:])
	})
}
