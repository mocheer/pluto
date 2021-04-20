package fn

import (
	"regexp"
	"strings"

	"github.com/mocheer/pluto/reg"
)

// ToCamelCase
// fn.ToCamelCase("camel-case") == "camelCase"
func ToCamelCase(val string) string {
	return regexp.MustCompile(reg.CamelCase).ReplaceAllStringFunc(val, strings.ToUpper)
}
