package dsfileinfo

import (
	"github.com/floyernick/fleep-go"
)

func GetInfoByBytes(bs []byte) (fleep.Info, error) {
	return fleep.GetInfo(bs)
}
