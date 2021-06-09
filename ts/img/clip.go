package img

import "image"

// Clip 剪切
func (p *Picture) Clip(x0, y0, width, height int) *Picture {
	return &Picture{Image: Clip(p.Image, x0, y0, width, height), Type: p.Type}
}

// Clip 剪切
func Clip(target image.Image, x0, y0, width, height int) image.Image {
	return target.(*image.RGBA).SubImage(image.Rect(x0, y0, x0+width, y0+height))
}
