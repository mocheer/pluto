package mh

import (
	"github.com/gomarkdown/markdown"
	"github.com/mocheer/pluto/pkg/fn"
)

type MH struct {
	text []byte
}

func New(text string) *MH {
	return &MH{
		text: fn.StringToBytes(text),
	}
}

// HTML 将一个mh对象转成html文本输出
func (m MH) HTML() string {
	return fn.BytesToString(markdown.ToHTML(m.text, nil, nil))
}
