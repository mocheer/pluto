package fn

import (
	"strings"
)

// SplitWords
// 以连续的空白字符为分隔符，将s切分成多个子串，结果中不包含空白字符本身。
//空白字符有：\\t, \\n, \\v, \\f, \\r, ’ ‘, U+0085 (NEL), U+00A0 (NBSP) 。
//如果 s 中只包含空白字符，则返回一个空列表\n  
func SplitWords(s string) []string {
	return strings.Fields(s)
}
