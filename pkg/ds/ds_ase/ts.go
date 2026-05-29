package ds_ase

// Aseprite 是aseprite的结构体
type Aseprite struct {
	Frames []FrameInfos `json:"frames"`
	Meta   Meta         `json:"meta"`
}

// FrameInfos 是帧的信息结构体
type FrameInfos struct {
	Filename         string           `json:"filename"`         // 文件名
	Frame            FrameRect        `json:"frame"`            // 帧的坐标,子图在精灵表中的位置和尺寸
	Rotated          bool             `json:"rotated"`          // 是否旋转90度
	Trimmed          bool             `json:"trimmed"`          // 是否裁剪
	SpriteSourceSize SpriteSourceSize `json:"spriteSourceSize"` // 精灵源大小,子图在原始画布上的裁剪后位置和尺寸, 仅在裁剪时有效
	SourceSize       Size             `json:"sourceSize"`       // 原始图的尺寸,不包含旋转和裁剪后的大小
	Duration         int              `json:"duration"`         // 帧的持续时间
}

// FrameRect 是帧的坐标结构体
type FrameRect struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

// SpriteSourceSize 是精灵源大小的结构体
type SpriteSourceSize struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

// Size 是大小的结构体
type Size struct {
	W int `json:"w"`
	H int `json:"h"`
}

// FrameTags 是帧标签的结构体
type FrameTags struct {
	Name      string `json:"name"`
	From      int    `json:"from"`
	To        int    `json:"to"`
	Direction string `json:"direction"`
}

// Meta 是元数据的结构体
type Meta struct {
	App       string      `json:"app"`
	Version   string      `json:"version"`
	Image     string      `json:"image"`
	Format    string      `json:"format"`
	Size      Size        `json:"size"`
	Scale     string      `json:"scale"`
	FrameTags []FrameTags `json:"frameTags"`
}
