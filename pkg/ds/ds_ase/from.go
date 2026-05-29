package ds_ase

import (
	"fmt"
	"image"
	"image/draw"
)

// FromImages
// 创建纹理集并返回一张大图和Aseprite数据
// 矩形装箱问题：Rectangle Packing
// 算法有：首次适应递减高度算法（First Fit Decreasing Height, FFDH），最佳适应算法（Best Fit），以及一些更复杂的算法如MaxRects算法
// https://github.com/InfinityTools/go-binpack2d 参考C实现的MaxRectsBinPack的算法，年代较久
// https://github.com/depp/skelly64/tree/main/lib/rectpack 任天堂的算法
// https://github.com/lewisgibson/go-binpack MaxRects算法，空间利用率极高，但性能小于skyline
// https://github.com/pekim/skyline 用于打包2D矩形的天际线算法。 空间利用率只是中等
// duration: 每帧的持续时间（毫秒），所有帧使用相同的时长
// tags: 动画标签，用于将帧分组为不同的动画剪辑（如 idle、run 等）
// 创建动画：3帧 idle，5帧 run，每帧 100ms
//
//	tags := []FrameTags{
//	    {Name: "idle", From: 0, To: 2, Direction: "forward"},
//	    {Name: "run",  From: 3, To: 7, Direction: "forward"},
//	}
//
// ase, atlas := FromImages(frames, 100, tags)
func FromImages(images []image.Image, duration int, tags []FrameTags) (*Aseprite, image.Image) {
	if len(images) == 0 {
		return &Aseprite{
			Meta: Meta{App: "aseprite", Version: "1.0", Format: "RGBA8888"},
		}, image.NewRGBA(image.Rect(0, 0, 0, 0))
	}

	// 计算大图尺寸 (简单横向堆叠)
	totalWidth := 0
	maxHeight := 0
	for _, img := range images {
		bounds := img.Bounds()
		totalWidth += bounds.Dx()
		if bounds.Dy() > maxHeight {
			maxHeight = bounds.Dy()
		}
	}

	// 创建空白大图
	atlas := image.NewRGBA(image.Rect(0, 0, totalWidth, maxHeight))

	ase := &Aseprite{
		Meta: Meta{
			App:       "aseprite",
			Version:   "1.0",
			Format:    "RGBA8888",
			Size:      Size{W: totalWidth, H: maxHeight},
			FrameTags: tags,
		},
	}

	xOffset := 0
	for i, img := range images {
		bounds := img.Bounds()
		width := bounds.Dx()
		height := bounds.Dy()

		// 绘制到纹理集
		draw.Draw(atlas, image.Rect(xOffset, 0, xOffset+width, height),
			img, bounds.Min, draw.Src)

		frame := FrameInfos{
			Filename:         fmt.Sprintf("frame_%d", i),
			Frame:            FrameRect{X: xOffset, Y: 0, W: width, H: height},
			Rotated:          false,
			Trimmed:          false,
			SpriteSourceSize: SpriteSourceSize{X: 0, Y: 0, W: width, H: height},
			SourceSize:       Size{W: width, H: height},
			Duration:         duration,
		}
		ase.Frames = append(ase.Frames, frame)
		xOffset += width
	}

	return ase, atlas
}
