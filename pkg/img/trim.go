package img

// import (
// 	"image"
// )
// Trim 裁剪图片, 移除图片每一行和每一列的空白区域，修改画布大小，获取最小的图片区域
// func Trim(img *image.Image) *image.Image {
// 	// 计算图片的空白区域
// 	bounds := img.Bounds()
// 	if bounds.Min.X == bounds.Max.X || bounds.Min.Y == bounds.Max.Y {
// 		return img
// 	}

// 	minX := bounds.Min.X -1
// 	minY := bounds.Min.Y -1	
// 	maxX := bounds.Max.X +1	
// 	maxY := bounds.Max.Y +1
// 	// 查找最小的非空白像素
// 	for x := bounds.Min.X; x < bounds.Max.X; x++ {
// 		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
// 			c := img.At(x, y)
// 			if c.A > 0 {
// 				minX = x
// 				break
// 			}
// 		}
// 		// 找到了minX，跳出循环
// 		if minX == x {
// 			break
// 		}
// 	}
// 	// 查找最大的非空白像素
// 	for x := bounds.Max.X; x > bounds.Min.X; x-- {
// 		for y := bounds.Max.Y; y > bounds.Min.Y; y-- {
// 			c := img.At(x, y)
// 			if c.A > 0 {
// 				maxX = x
// 				break
// 			}
// 		}
// 		// 找到了maxX，跳出循环
// 		if maxX == x {
// 			break
// 		}
// 	}
// 	// 
// 	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
// 		for x := bounds.Min.X; x < bounds.Max.X; x++ {
// 			c := img.At(x, y)
// 			if c.A > 0 {
// 				minY = y
// 				break
// 			}
// 		}
// 		// 找到了minY，跳出循环
// 		if minY == y {
// 			break
// 		}
// 	}
// 	// 查找最大的非空白像素
// 	for y := bounds.Max.Y; y > bounds.Min.Y; y-- {
// 		for x := bounds.Max.X; x > bounds.Min.X; x-- {
// 			c := img.At(x, y)
// 			if c.A > 0 {
// 				maxY = y
// 				break
// 			}
// 		}
// 		// 找到了maxY，跳出循环
// 		if maxY == y {
// 			break
// 		}
// 	}

// 	// 裁剪图片
// 	img = img.SubImage(image.Rect(minX, minY, maxX, maxY))
// 	return img
// }