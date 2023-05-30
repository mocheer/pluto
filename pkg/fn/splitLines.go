package fn

import (
	"bufio"
	"regexp"
	"strings"
)

// SplitLines 按行读取字符串
func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

// SplitLines2 根据换行符分隔字符串
func SplitLines2(s string) []string {
	zp := regexp.MustCompile(`[\t\n\f\r]`)
	return zp.Split(s, -1)
}
