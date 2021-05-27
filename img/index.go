package img

import (
	"bytes"
	"image"
	"io"
	"os"
)

type Picture struct {
	Image image.Image
	Type  string
}

// FromReader
func FromReader(r io.Reader) (*Picture, error) {
	i, imageType, err := image.Decode(r)
	return &Picture{Image: i, Type: imageType}, err
}

// FromBytes
func FromBytes(bs []byte) (*Picture, error) {
	return FromReader(bytes.NewBuffer(bs))
}

// FromFile
func FromFile(path string) (*Picture, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return FromReader(f)
}
