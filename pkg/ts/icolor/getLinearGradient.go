package icolor

import (
	"image/color"
)

// getLinearGradient LerpColor calculates the interpolated color between two colors based on a value t, where t=0
// returns color1 and t=1 returns color2.
func GetLinearGradient(color1, color2 color.Color, t float64) color.RGBA64 {
	r1, g1, b1, a1 := color1.RGBA()
	r2, g2, b2, a2 := color1.RGBA()

	// Interpolate each color component
	r := uint8(float64(r1)*(1-t) + float64(r2)*t)
	g := uint8(float64(g1)*(1-t) + float64(g2)*t)
	b := uint8(float64(b1)*(1-t) + float64(b2)*t)
	a := uint8(float64(a1)*(1-t) + float64(a2)*t)

	return color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)}
}
