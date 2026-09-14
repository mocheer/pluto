package img

import (
	"bytes"
	"image"
	"io"
	"os"
)

type Img struct {
	Image image.Image
}

// FromFile 从文件中读取数据实例化Picture对象
func FromFile(path string) (*Img, string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, "", err
	}
	return FromReader(f)
}

// FromBytes 从bytes数据中实例化Picture对象
func FromBytes(bs []byte) (*Img, string, error) {
	return FromReader(bytes.NewBuffer(bs))
}

// FromReader
func FromReader(r io.Reader) (*Img, string, error) {
	i, format, err := image.Decode(r)
	return &Img{Image: i}, format, err
}

// FromImage
func FromImage(ig image.Image) *Img {
	return &Img{Image: ig}
}
