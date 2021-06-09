package img

import (
	"image"

	"github.com/nfnt/resize"
)

// Resize 重置大小
func (p *Picture) Resize(width, height uint) *Picture {
	return &Picture{Image: Resize(p.Image, width, height), Type: p.Type}
}

// Thumbnail 缩略图
func (p *Picture) Thumbnail(width, height uint) *Picture {
	return &Picture{Image: Thumbnail(p.Image, width, height), Type: p.Type}
}

// Resize 当宽度或者高度为0时，保持比例
func Resize(target image.Image, width, height uint) image.Image {
	return resize.Resize(width, height, target, resize.Lanczos3)
}

// Thumbnail
func Thumbnail(target image.Image, maxWidth, maxHeight uint) image.Image {
	return resize.Thumbnail(maxWidth, maxHeight, target, resize.Lanczos3)
}
