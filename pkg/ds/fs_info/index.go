package fs_info

import (
	"github.com/floyernick/fleep-go"
)

// GetInfoByBytes
func GetInfoByBytes(bs []byte) (fleep.Info, error) {
	return fleep.GetInfo(bs)
}
