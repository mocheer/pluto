package fn

import (
	"regexp"
	"strings"
)

// ToCamelCase
// ToCamelCase("camel-case") == "camelCase"
// ToCamelCase("camel_case") == "camelCase"
func ToCamelCase(val string) string {
	// 匹配短连接字符`-`，常用于转成驼峰大写
	return regexp.MustCompile(`[-_]+(.)?`).ReplaceAllStringFunc(val, func(s string) string {
		return strings.ToUpper(s[1:])
	})
}

func ToCamelCase2(val string) string {
	return strings.ToUpper(val[:1]) + ToCamelCase(val[1:])
}
