package img

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
)

// ToBytes 将图片对象转成bytes字节流
func (m *Img) ToBytes() (bs []byte, err error) {
	return ToBytes(m.Image, m.Type)
}

// MustToBytes
func (m *Img) MustToBytes() []byte {
	bs, err := ToBytes(m.Image, m.Type)
	if err != nil {
		panic(err)
	}
	return bs
}

// ToBytes 将图片对象转成bytes字节流
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
