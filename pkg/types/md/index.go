package md

import (
	"bytes"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/mocheer/pluto/pkg/fn"
)

type MD struct {
	Buffer bytes.Buffer
}

func New() *MD {
	return new(MD)
}

func (m *MD) Write(p []byte) {
	m.Buffer.Write(p)
}

func (m *MD) WriteString(s string) {
	m.Buffer.WriteString(s)
}

func (m *MD) H(index int, title string) {
	m.WriteString(strings.Repeat("#", index) + " " + title + "\n")
}

func (m *MD) H1(title string) {
	m.H(1, title)
}

func (m *MD) H2(title string) {
	m.H(2, title)
}

func (m *MD) H3(title string) {
	m.H(3, title)
}

func (m *MD) Text(text string) {
	m.WriteString(text)
}

// Code 代码块
func (m *MD) Code(code string, lang string) {
	m.WriteString("```" + lang)
	m.LineBreak()
	m.WriteString(code)
	m.LineBreak()
	m.WriteString("```")
	m.LineBreak()
}

// LineBreak 换行
func (m *MD) LineBreak() {
	m.WriteString("\n")
}

// Paragraph 段落
func (m *MD) Paragraph(text string) {
	m.WriteString(text)
	m.LineBreak()
}

// Image 图片
func (m *MD) Image(url string, alt string) {
	m.WriteString("![" + alt + "](" + url + ")")
	m.LineBreak()
}

// Link 链接
func (m *MD) Link(url string, text string) {
	m.WriteString("[" + text + "](" + url + ")")
	m.LineBreak()
}

// Table 表格
func (m *MD) Table(headers, rows []string) {
	m.WriteString(strings.Join(headers, "|") + "\n")
	m.WriteString(strings.Repeat("-", len(headers)*3-1) + "\n")
	m.WriteString(strings.Join(rows, "|") + "\n")
}

// ToHTML 转成html文本输出
// 弃用，TODO修改为github.com/yuin/goldmark
func (m *MD) ToHTML() string {
	return fn.BytesToString(markdown.ToHTML(m.Buffer.Bytes(), nil, nil))
}
