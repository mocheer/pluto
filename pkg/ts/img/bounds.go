package img

import "image"

func Bounds(m image.Image) (int, int) {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	return dx, dy
}
