package ds_text

import (
	"os"

	"github.com/mocheer/pluto/ds"
	"github.com/mocheer/pluto/fn"
)

//
func Read(fileName string) (string, error) {
	bs, err := os.ReadFile(fileName)
	return fn.B2S(bs), err
}

// MustRead 读取文本文件，当发生错误的时候直接panic
func MustRead(fileName string) string {
	return fn.B2S(ds.MustRead(fileName))
}
