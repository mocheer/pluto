package img

import "image"

// Bounds 获取图片范围
func Bounds(m image.Image) (int, int) {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	return dx, dy
}
