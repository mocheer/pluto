package mh

import (
	"github.com/gomarkdown/markdown"
	"github.com/mocheer/pluto/fn"
)

type MH struct {
	text []byte
}

//
func New(text string) *MH {
	return &MH{
		text: fn.S2B(text),
	}
}

// HTML 将一个mh对象转成html文本输出
func (m MH) HTML() string {
	return fn.B2S(markdown.ToHTML(m.text, nil, nil))
}
