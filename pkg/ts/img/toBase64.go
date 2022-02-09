package img

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"

	js "github.com/mocheer/pluto/pkg/jsg"
)

func (p *Img) MustToBase64() (data string) {
	data, err := ToBase64(p.Image, p.Type)
	if err != nil {
		panic(err)
	}
	return
}

func (p *Img) ToBase64() (data string, err error) {
	return ToBase64(p.Image, p.Type)
}

// ToJPEGBase64 将image转成各种图片编码的base64字符串
func ToBase64(target image.Image, imageType string) (data string, err error) {
	buf := new(bytes.Buffer)
	switch imageType {
	case JPEG:
		err = jpeg.Encode(buf, target, nil)
	case PNG:
		err = png.Encode(buf, target)
	case GIF:
		err = gif.Encode(buf, target, &gif.Options{})
	}
	if err != nil {
		return
	}
	data = js.Btoa(buf.String())
	return
}
