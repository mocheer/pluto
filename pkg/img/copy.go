package img

import (
	"image"
	"image/draw"
)

// Copy
func Copy(src image.Image) draw.Image {
	srcBounds := src.Bounds()
	copy := image.NewRGBA(srcBounds)
	draw.Draw(copy, srcBounds, src, image.Point{}, draw.Over)
	return copy
}
