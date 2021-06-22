package fn

import (
	"fmt"
	"strconv"
	"strings"
)

// Unicode2ZH 将unicode编码转成中文
func Unicode2ZH(from string) string {
	textQuoted := strconv.QuoteToASCII(from)
	textUnquoted := textQuoted[1 : len(textQuoted)-1]

	sUnicodev := strings.Split(textUnquoted, "\\u")
	var context string
	for _, v := range sUnicodev {
		if len(v) < 1 {
			continue
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			panic(err)
		}
		context += fmt.Sprintf("%c", temp)
	}
	return context
}
