package atlas

import "image"

// Atlas 对应 .atlas JSON 文件的顶层结构
// 包含了所有子图的图像数据和帧信息。
// 每个子图的图像数据存储在 Images 字段中，每个帧的图像数据存储在 Frames 字段中。
// 每个子图的图像数据都是一个 image.Image 实例，每个帧的图像数据都是一个 SpriteRectInfo 结构，描述了帧的位置和尺寸。
// 每个子图的图像数据都是一个 image.Image 实例，描述了子图的图像内容。
type Atlas struct {
	Images []image.Image    `json:"images"`
	Frames map[string]Frame `json:"frames"`
}

// Frame 对应 .atlas JSON 中每个帧的结构
// 每个帧包含一个 SpriteRectInfo 结构，描述了帧的位置和尺寸，以及一个 image.Image 实例，描述了帧的图像内容。
type Frame struct {
	Frame    SpriteRectInfo `json:"frame"`
	SubImage image.Image    `json:"subImage"`
}

// subImager 是一个接口，用于描述一个可以裁剪图像的类型。
type subImager interface {
	SubImage(r image.Rectangle) image.Image
}
