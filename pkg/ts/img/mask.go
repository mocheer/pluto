package img

import (
	"image"
	"image/color"
	"image/draw"
)

// MaskImage 获取蒙版遮罩镂空的背景图（裁剪后的区域）
func MaskImage(src, mask image.Image, copyPoint image.Point) (draw.Image, error) {
	srcBounds := src.Bounds()
	maskBounds := mask.Bounds()

	white := image.Uniform{
		C: color.RGBA{255, 255, 255, 255},
	}
	whiteImg := image.NewRGBA(maskBounds)
	draw.Draw(whiteImg, maskBounds, &white, image.Point{}, draw.Over)

	copy := copySrc(src)
	draw.DrawMask(copy, srcBounds.Add(copyPoint), whiteImg, image.Point{}, mask, maskBounds.Min, draw.Over)

	return copy, nil
}

// PieceImage 获取蒙版遮罩本身的空白区域（裁剪下来的区域）
func PieceImage(src, mask image.Image, copyPoint image.Point) (draw.Image, error) {
	maskBounds := mask.Bounds()
	// Create a new image with mask bounds for final move block
	copy := copySrc(mask)
	// Get the part image in src image with the mask-bounds
	draw.DrawMask(copy, maskBounds, src, copyPoint, mask, maskBounds.Min, draw.Over)

	return copy, nil
}
