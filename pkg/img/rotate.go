package img

import (
	"image"
	"image/color"

	"github.com/disintegration/imaging"
)

// Rotate
func (p *Img) Rotate(angle float64) *Img {
	return &Img{Image: Rotate(p.Image, angle)}
}

// Rotate
func (p *Img) Rotate90() *Img {
	return &Img{Image: Rotate90(p.Image)}
}

// Rotate
func (p *Img) Rotate270() *Img {
	return &Img{Image: Rotate270(p.Image)}
}

// Rotate
func Rotate(target image.Image, angle float64) image.Image {
	return imaging.Rotate(target, angle, color.RGBA64{})
}

// Rotate
func Rotate90(target image.Image) image.Image {
	return imaging.Rotate90(target)
}

// Rotate
func Rotate270(target image.Image) image.Image {
	return imaging.Rotate270(target)
}
