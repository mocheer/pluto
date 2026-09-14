package img

import (
	"image"
	"image/color"
)

// SetColorTransparent
// c 一般为黑色背景的RGB值
func SetColorTransparent(source image.Image, c color.RGBA) image.Image {
	bounds := source.Bounds()
	target := image.NewRGBA(bounds)
	// 遍历原始图像的每个像素并处理背景色
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := source.At(x, y)
			r, g, b, _ := pixel.RGBA()
			if r == uint32(c.R) && g == uint32(c.G) && b == uint32(c.B) {
				target.SetRGBA(x, y, color.RGBA{0, 0, 0, 0})
			} else {
				target.Set(x, y, pixel) // 其他颜色保持不变
			}
		}
	}
	return target
}

// SetColorTransparentWithBlack
func SetColorTransparentWithBlack(source image.Image) image.Image {
	return SetColorTransparent(source, color.RGBA{0, 0, 0, 255})
}
