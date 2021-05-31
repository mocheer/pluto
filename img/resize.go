package img

import (
	"image"

	"github.com/nfnt/resize"
)

// Resize
func (p *Picture) Resize(width, height uint) *Picture {
	return &Picture{Image: Resize(p.Image, width, height), Type: p.Type}
}

// Resize 当宽度或者高度为0时，保持比例
func Resize(target image.Image, width, height uint) image.Image {
	return resize.Resize(width, height, target, resize.Lanczos3)
}
