package fn

import (
	"fmt"
	"regexp"

	"github.com/mocheer/pluto/reg"
)

// FmtString
// @param str
// @param data
func FmtString(src string, data map[string]interface{}) string {
	return regexp.MustCompile(reg.Brace).ReplaceAllStringFunc(src, func(key string) string {
		// 这里的key包含括号
		val := data[key[1:len(key)-1]]
		if val == nil {
			return ""
		}
		return fmt.Sprint(val)
	})
}
