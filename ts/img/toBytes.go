package img

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
)

func (p *Picture) ToBytes() (bs []byte, err error) {
	return ToBytes(p.Image, p.Type)
}

func ToBytes(target image.Image, imageType string) (bs []byte, err error) {
	buf := new(bytes.Buffer)
	switch imageType {
	case JPEG:
		err = jpeg.Encode(buf, target, &jpeg.Options{})
	case PNG:
		err = png.Encode(buf, target)
	case GIF:
		err = gif.Encode(buf, target, &gif.Options{})
	}
	if err != nil {
		return
	}
	bs = buf.Bytes()
	return
}
