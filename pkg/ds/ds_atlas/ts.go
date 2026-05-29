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

// atlasConfig 对应 .atlas JSON 文件的顶层结构
type atlasConfig struct {
	Meta   meta             `json:"meta"`   // 元数据信息
	Frames map[string]frame `json:"frames"` // 帧信息
}

// Frame 对应 .atlas JSON 中每个帧的结构
// 每个帧包含一个 SpriteRectInfo 结构，描述了帧的位置和尺寸，以及一个 image.Image 实例，描述了帧的图像内容。
type Frame struct {
	Frame    SpriteRectInfo `json:"frame"`
	SubImage image.Image    `json:"subImage"`
}

// meta 对应 .atlas JSON 中的元数据信息
type meta struct {
	Image  string  `json:"image"`
	Format string  `json:"format"`
	Scale  float64 `json:"scale"` // 可选参数：缩放比例
}

// frame 对应 .atlas JSON 中每个帧的结构（内部解析用）
// 在 LayaAir 的 .atlas 文件中，"frames" 字段里的每一个条目描述的是一张独立的子图（或称精灵图片、Sprite/Sub‑image），而不是动画中的“帧”序列。
type frame struct {
	Frame            SpriteRectInfo   `json:"frame"`            // 帧信息
	Rotated          bool             `json:"rotated"`          // 是否旋转
	Trimmed          bool             `json:"trimmed"`          // 是否裁剪
	SpriteSourceSize SpriteSourceSize `json:"spriteSourceSize"` // 当 trimmed 为 true 时 描述“剪裁”后图片在原始未经剪裁的画布中的位置和尺寸。
	SourceSize       Size             `json:"sourceSize"`       // 当 trimmed 为 true 时 记录原始图片未经任何剪裁时的原始宽度和高度。
}

type Size struct {
	W int `json:"w"`
	H int `json:"h"`
}

type SpriteSourceSize struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

type SpriteRectInfo struct {
	Idx int `json:"idx"` // 帧索引，不一定存在
	X   int `json:"x"`
	Y   int `json:"y"`
	W   int `json:"w"`
	H   int `json:"h"`
}

type subImager interface {
	SubImage(r image.Rectangle) image.Image
}
