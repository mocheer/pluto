package fn

import "strings"

// replaceChars 将字符串中的字符根据映射关系替换
// 定义映射关系
// replacements := map[rune]rune{
// 	'a': 'b',
// 	'b': 'c',
// }
// // 原始字符串
// originalStr := "I am a student."
// // 替换字符串中的字符
// newStr := replaceChars(originalStr, replacements)
// fmt.Println("Original:", originalStr)
// fmt.Println("Modified:", newStr)
func ReplaceChars(s string, replacements map[rune]rune) string {
	var builder strings.Builder
	for _, r := range s {
		if replacement, ok := replacements[r]; ok {
			builder.WriteRune(replacement)
		} else {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}
