package md

import (
	"github.com/gomarkdown/markdown"
	"github.com/mocheer/pluto/pkg/fn"
)

type MD []byte

// HTML 将一个mh对象转成html文本输出
// 弃用，TODO修改为github.com/yuin/goldmark
func (m MD) HTML() string {
	return fn.BytesToString(markdown.ToHTML(m, nil, nil))
}

// func (m MD) HTML2() string {
// 	return fn.BytesToString(blackfriday.Run(m))
// }
