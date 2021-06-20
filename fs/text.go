package fs

import (
	"github.com/mocheer/pluto/fn"
)

// MustReadText 读取文本文件，当发生错误的时候直接panic
func MustReadText(fileName string) string {
	return fn.B2S(MustRead(fileName))
}
