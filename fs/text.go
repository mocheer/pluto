package fs

import (
	"github.com/mocheer/pluto/fn"
)

// MustReadText 读取文本文件，当发生错误的时候直接panic
func MustReadText(path string) string {
	return fn.Bytes2String(MustReadFile(path))
}
