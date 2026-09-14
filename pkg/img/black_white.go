package img

import (
	"image"
	"image/color"
	"log"
	"slices"
)

// 用于通过透明的线框图制作黑白图

// 8个邻域方向（顺时针顺序）
var directions = [8][2]int{
	{0, -1},  // 上
	{1, -1},  // 右上
	{1, 0},   // 右
	{1, 1},   // 右下
	{0, 1},   // 下
	{-1, 1},  // 左下
	{-1, 0},  // 左
	{-1, -1}, // 左上
}

func Edge(img image.Image) *image.RGBA {

	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	processedImg := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			colorValue := img.At(x, y)
			processedImg.Set(x, y, colorValue)
		}
	}
	// 创建访问标记数组
	visited := make([][]bool, width)
	for i := range visited {
		visited[i] = make([]bool, height)
	}

	edgePoints := []image.Point{{0, 0}}
	whiteRGBA := color.RGBA64{65535, 65535, 65535, 65535}
	blackRGBA := color.RGBA64{0, 0, 0, 65535}
	processedImg.SetRGBA64(0, 0, whiteRGBA)
	for len(edgePoints) > 0 {
		start := edgePoints[0]
		for _, p := range directions {
			x := start.X + p[0]
			y := start.Y + p[1]

			if x >= 0 && x < width && y >= 0 && y < height {
				if visited[x][y] {
					continue
				}
				visited[x][y] = true
				_, _, _, a := img.At(x, y).RGBA()

				if a < 1 {
					edgePoints = append(edgePoints, image.Point{x, y})
				}
				// 这里会将边界1px范围也设置为白色
				processedImg.SetRGBA64(x, y, whiteRGBA)
			}
		}
		edgePoints = slices.Delete(edgePoints, 0, 1)
	}

	log.Println(processedImg.At(0, 0).RGBA())
	if true {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				r, g, b, a := processedImg.At(x, y).RGBA()
				if r != 65535 || g != 65535 || b != 65535 || a != 65535 {
					processedImg.SetRGBA64(x, y, blackRGBA)
				}
			}
		}
	}

	return processedImg

}
