package img

import (
	"image"
	"image/color"
)

// Each
func Each(m image.Image, callback func(color.Color, int, int)) {
	dx, dy := Bounds(m)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			callback(m.At(i, j), i, j)
		}
	}
}
