package icolor

import "image/color"

func GetBrightness(color color.Color) float64 {
	r, g, b, _ := color.RGBA()
	return (0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 65535
}
