package img

import (
	"image"
	"image/draw"
)

func copySrc(src image.Image) draw.Image {
	srcBounds := src.Bounds()
	copy := image.NewRGBA(srcBounds)
	draw.Draw(copy, srcBounds, src, image.Point{}, draw.Over)
	return copy
}
