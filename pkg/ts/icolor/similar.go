package icolor

import (
	"image/color"
	"math"
)

// Similar 给定差值，判断两个颜色是否相似，可用于等值面图例识别等（有些图片失真或者有渐变导致颜色不一定完全根据图例绘制）
// 透明度不参与计算
func Similar(c1, c2 color.Color, tolerance float64) bool {
	r1, g1, b1, _ := c1.RGBA()
	r2, g2, b2, _ := c2.RGBA()
	dr := math.Abs(float64(r1) - float64(r2))
	dg := math.Abs(float64(g1) - float64(g2))
	db := math.Abs(float64(b1) - float64(b2))
	return (dr < tolerance) && (dg < tolerance) && (db < tolerance)
}

// Similar2
func Similar2(c1, c2 color.Color, tolerance float64) bool {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	dr := math.Abs(float64(r1) - float64(r2))
	dg := math.Abs(float64(g1) - float64(g2))
	db := math.Abs(float64(b1) - float64(b2))
	da := math.Abs(float64(a1) - float64(a2))
	return dr < tolerance && dg < tolerance && db < tolerance && da < tolerance
}
