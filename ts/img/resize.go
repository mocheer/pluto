package img

import (
	"image"

	"github.com/disintegration/imaging"
)

// Resize 重置大小
func (p *Picture) Resize(width, height int) *Picture {
	return &Picture{Image: Resize(p.Image, width, height), Type: p.Type}
}

// Resize 当宽度或者高度为0时，保持比例
func Resize(target image.Image, width, height int) image.Image {
	return imaging.Resize(target, width, height, imaging.Lanczos)
}
