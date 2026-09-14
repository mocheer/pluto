package img

import (
	"image"

	"github.com/mocheer/pluto/pkg/ts/icolor"
)

// MergeTileImage
// 用于合并有黑色背景的mbtiles的图片
func MergeTileImage(source image.Image, source2 image.Image) image.Image {
	bounds := source.Bounds()
	target := image.NewRGBA(bounds)
	// 遍历原始图像的每个像素并处理背景色
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel1 := source.At(x, y)
			pixel2 := source2.At(x, y)
			if icolor.GetBrightness(pixel1) > icolor.GetBrightness(pixel2) {
				target.Set(x, y, pixel1)
			} else {
				target.Set(x, y, pixel2)
			}
		}
	}
	return target

}
