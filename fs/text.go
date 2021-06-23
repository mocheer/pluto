package fs

import (
	"github.com/mocheer/pluto/fn"
)

//
func ReadText(fileName string) (string, error) {
	bs, err := Read(fileName)
	return fn.B2S(bs), err
}

// MustReadText 读取文本文件，当发生错误的时候直接panic
func MustReadText(fileName string) string {
	return fn.B2S(MustRead(fileName))
}
