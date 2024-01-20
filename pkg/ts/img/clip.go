package img

import (
	"image"
)

// 可以用draw包进行截取，目前用的sub

// Clip 剪切
func (m *Img) Clip(x0, y0, width, height int) *Img {
	return &Img{Image: Clip(m.Image, x0, y0, width, height)}
}

// Clip 剪切
func (m *Img) ClipPNG(x0, y0, width, height int) *Img {
	return &Img{Image: ClipPNG(m.Image, x0, y0, width, height)}
}

// Clip 剪切
func Clip(target image.Image, x0, y0, width, height int) image.Image {
	return target.(*image.RGBA).SubImage(image.Rect(x0, y0, x0+width, y0+height))
}

// ClipPNG 裁剪
func ClipPNG(target image.Image, x0, y0, width, height int) image.Image {
	return target.(*image.NRGBA).SubImage(image.Rect(x0, y0, x0+width, y0+height))
}

// ClipImage 剪切
func ClipImage(target image.Image, x0, y0, width, height int) image.Image {
	return nil
}
