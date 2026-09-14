//go:build linux

package ds_webp

import (
	"bytes"
	"image"
	"io"

	"github.com/mocheer/pluto/pkg/img"
)

// func RegisterFormat() {
// 	image.RegisterFormat("webp", "RIFF????WEBPVP8", webp.Decode, webp.DecodeConfig)
// }

func Decode(r io.Reader) (image.Image, error) {
	// return webp.Decode(r)
	return nil, nil
}

// FromImageFile
func FromImageFile(filename string) (*bytes.Buffer, error) {
	i, _, err := img.FromFile(filename)
	if err != nil {
		return nil, err
	}
	return FromImage(i.Image)
}

// FromImage
func FromImage(m image.Image) (*bytes.Buffer, error) {
	// var bs bytes.Buffer
	// err := webp.Encode(&bs, m, &webp.Options{Lossless: false, Quality: 75})
	// return &bs, err
	return nil, nil
}

// Save
func Save(filename string, m image.Image) error {
	// return webp.Save(filename, m, &webp.Options{Lossless: false, Quality: 75})
	return nil
}
