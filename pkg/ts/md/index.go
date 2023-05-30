package md

import (
	"github.com/gomarkdown/markdown"
	"github.com/mocheer/pluto/pkg/fn"
)

type MD []byte

// HTML 将一个mh对象转成html文本输出
func (m MD) HTML() string {
	return fn.BytesToString(markdown.ToHTML(m, nil, nil))
}
