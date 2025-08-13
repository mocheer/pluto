package img

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
)

// ToBytes 将图片对象转成bytes字节流
func (m *Img) ToBytes() (bs []byte, err error) {
	return ToBytes(m.Image, JPG)
}

// MustToBytes
func (m *Img) MustToBytes(format string) []byte {
	bs, err := ToBytes(m.Image, format)
	if err != nil {
		panic(err)
	}
	return bs
}

// ToBytes 将图片对象转成bytes字节流
func ToBytes(target image.Image, imageType string) (bs []byte, err error) {
	var buf bytes.Buffer
	switch imageType {
	case JPG, JPEG:
		// 默认的压缩质量为75，损失的图片细节太多
		err = jpeg.Encode(&buf, target, &jpeg.Options{Quality: 90})
	case PNG:
		err = png.Encode(&buf, target)
	case GIF:
		err = gif.Encode(&buf, target, &gif.Options{})
	// case WEBP:
	default:
		err = fmt.Errorf("格式[%s]不支持", imageType)
	}
	if err == nil {
		bs = buf.Bytes()
	}
	return
}

func ToJpegBytes(target image.Image, options *jpeg.Options) (bs []byte, err error) {
	var buf bytes.Buffer
	err = jpeg.Encode(&buf, target, options)
	if err == nil {
		bs = buf.Bytes()
	}
	return
}
