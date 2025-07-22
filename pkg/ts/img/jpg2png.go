package img

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)

// TODO
// 获取图像四个角的颜色作为背景候选
// corners := []image.Point{bounds.Min, bounds.Max, {bounds.Min.X, bounds.Max.Y}, {bounds.Max.X, bounds.Min.Y}}
// 统计出现频率最高的颜色
//
// 使用github.com/disintegration/imaging库进行高斯模糊
// blurredImg := imaging.Blur(transparentImg, 0.5)

// ConvertJPEGToTransparentPNG 将指定背景色的JPEG转换为透明PNG
// 参考 img.SetColorTransparent
// 参数：
//
//	inputPath: 输入JPEG文件路径
//	outputPath: 输出PNG文件路径
//	bgColor: 目标背景色（RGBA格式）
//	tolerance: 颜色容差值（0-255）
//
// 返回值：
//
//	error: 操作过程中发生的错误
func ConvertJPEGToTransparentPNG(inputPath, outputPath string, bgColor color.RGBA, tolerance uint8) error {
	// 1. 打开并解码JPEG
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	// 2. 创建透明图像
	bounds := img.Bounds()
	transparentImg := image.NewNRGBA(bounds)

	// 3. 处理像素透明度
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)

			if isBackground(r8, g8, b8, bgColor, tolerance) {
				transparentImg.Set(x, y, color.NRGBA{R: r8, G: g8, B: b8, A: 0})
			} else {
				transparentImg.Set(x, y, img.At(x, y))
			}
		}
	}

	// 4. 保存PNG
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return png.Encode(outFile, transparentImg)
}

// 判断颜色是否属于背景（未修改）
func isBackground(r, g, b uint8, target color.RGBA, tolerance uint8) bool {
	deltaR := abs(int(r) - int(target.R))
	deltaG := abs(int(g) - int(target.G))
	deltaB := abs(int(b) - int(target.B))
	return deltaR <= int(tolerance) && deltaG <= int(tolerance) && deltaB <= int(tolerance)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
