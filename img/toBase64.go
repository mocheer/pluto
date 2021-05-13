package img

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"

	"github.com/mocheer/pluto/js/window"
)

// ToJPEGBase64 将image转成各种图片编码的base64字符串
func ToBase64(target image.Image, imageType string) (data string, err error) {
	buf := new(bytes.Buffer)
	switch imageType {
	case "jpeg":
		err = jpeg.Encode(buf, target, nil)
	case "png":
		err = png.Encode(buf, target)
	case "gif":
		err = gif.Encode(buf, target, &gif.Options{})
	}
	if err != nil {
		return
	}
	data = window.Btoa(buf.String())
	return
}
