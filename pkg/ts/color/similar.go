package color

import (
	"image/color"
	"math"
)

// Similar 给定差值，判断两个颜色是否相似，可用于等值面图例识别等（有些图片失真或者有渐变导致颜色不一定完全根据图例绘制）
func Similar(c1, c2 color.RGBA, tolerance float64) bool {
	dr := int(c1.R) - int(c2.R)
	dg := int(c1.G) - int(c2.G)
	db := int(c1.B) - int(c2.B)
	da := int(c1.A) - int(c2.A)

	return math.Abs(float64(dr)) < tolerance && math.Abs(float64(dg)) < tolerance && math.Abs(float64(db)) < tolerance && math.Abs(float64(da)) < tolerance
}
