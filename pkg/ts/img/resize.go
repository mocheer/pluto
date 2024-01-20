package img

import (
	"image"

	"github.com/disintegration/imaging"
)

// Resize 重置大小
func (p *Img) Resize(width, height int) *Img {
	return &Img{Image: Resize(p.Image, width, height)}
}

// Resize 当宽度或者高度为0时，保持比例
func Resize(target image.Image, width, height int) image.Image {
	return imaging.Resize(target, width, height, imaging.Lanczos)
}
