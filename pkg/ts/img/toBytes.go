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
	case JPG:
		fallthrough
	case JPEG:
		// 默认的压缩质量为75，损失的图片细节太多
		err = jpeg.Encode(buf, target, &jpeg.Options{Quality: 92})
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
