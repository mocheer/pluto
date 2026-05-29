package ds_ase

// Aseprite 是aseprite的结构体
type Aseprite struct {
	Frames []Frames `json:"frames"`
	Meta   Meta     `json:"meta"`
}

// Frame 是帧的结构体
type Frame struct {
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

// Frames 是帧的结构体
type Frames struct {
	Filename         string           `json:"filename"`
	Frame            Frame            `json:"frame"`
	Rotated          bool             `json:"rotated"`
	Trimmed          bool             `json:"trimmed"`
	SpriteSourceSize SpriteSourceSize `json:"spriteSourceSize"`
	SourceSize       Size             `json:"sourceSize"`
	Duration         int              `json:"duration"`
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
