package img

import "image"

// Clip 剪切image
func Clip(target image.Image, x0, y0, width, height int) image.Image {
	return target.(*image.RGBA).SubImage(image.Rect(x0, y0, x0+width, y0+height))
}
