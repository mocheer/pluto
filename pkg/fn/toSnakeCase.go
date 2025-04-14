package fn

import (
	"regexp"
	"strings"
)

// ToSnakeCase
// ToSnakeCase("snakeCase") == "snake-case"
func ToSnakeCase(val string) string {
	// 匹配驼峰转成短连接符
	return ToSnakeCaseWithSplit(val, "-")
}

func ToSnakeCaseWithSplit(val string, splitStr string) string {
	return regexp.MustCompile(`([a-z])([A-Z])`).ReplaceAllStringFunc(val, func(s string) string {
		return s[:1] + splitStr + strings.ToLower(s[1:])
	})
}

func ToSnakeCaseWithSplit2(val string, splitStr string) string {
	return strings.ToLower(val[:1]) + ToSnakeCaseWithSplit(val[1:], splitStr)
}
