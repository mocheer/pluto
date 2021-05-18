package fn

import (
	"regexp"
	"strings"

	"github.com/mocheer/pluto/reg"
)

// ToCamelCase
func ToCamelCase(val string) string {
	return regexp.MustCompile(reg.CamelCase).ReplaceAllStringFunc(val, func(s string) string {
		return strings.ToUpper(s[1:])
	})
}
