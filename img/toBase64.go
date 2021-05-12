package img

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
)

// ToJPEGBase64 将image转成基于jpg编码的base64图片
func ToJPEGBase64(target image.Image) (string, error) {
	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, target, nil)
	if err != nil {
		return "", err
	}
	return BytesToBase64(buf.Bytes()), nil
}

// ToJPEGBase64 将image转成基于jpg编码的base64图片
func ToPNGBase64(target image.Image) (string, error) {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, target)
	if err != nil {
		return "", err
	}
	return BytesToBase64(buf.Bytes()), nil
}
