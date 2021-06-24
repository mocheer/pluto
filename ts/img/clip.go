package img

import "image"

// Clip 剪切
func (m *Img) Clip(x0, y0, width, height int) *Img {
	return &Img{Image: Clip(m.Image, x0, y0, width, height), Type: m.Type}
}

// Clip 剪切
func Clip(target image.Image, x0, y0, width, height int) image.Image {
	return target.(*image.RGBA).SubImage(image.Rect(x0, y0, x0+width, y0+height))
}
