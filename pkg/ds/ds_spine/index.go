package ds_spine

import (
	"encoding/json"
	"os"
)

// SkeletonMeta 骨骼元数据
type SkeletonMeta struct {
	Hash    string  `json:"hash,omitempty"`
	Spine   string  `json:"spine,omitempty"`
	X       float64 `json:"x,omitempty"`
	Y       float64 `json:"y,omitempty"`
	Width   float64 `json:"width,omitempty"`
	Height  float64 `json:"height,omitempty"`
	Images  string  `json:"images,omitempty"`
	Audio   string  `json:"audio,omitempty"`
	FPS     float64 `json:"fps,omitempty"`
	RealFPS float64 `json:"realFps,omitempty"`
}

// Bone 骨骼定义
type Bone struct {
	Name            string  `json:"name"`
	Parent          string  `json:"parent,omitempty"`
	X               float64 `json:"x,omitempty"`
	Y               float64 `json:"y,omitempty"`
	Rotation        float64 `json:"rotation,omitempty"`
	ScaleX          float64 `json:"scaleX,omitempty"`
	ScaleY          float64 `json:"scaleY,omitempty"`
	ShearX          float64 `json:"shearX,omitempty"`
	ShearY          float64 `json:"shearY,omitempty"`
	Length          float64 `json:"length,omitempty"`
	Color           string  `json:"color,omitempty"`
	InheritScale    bool    `json:"inheritScale,omitempty"`
	InheritRotation bool    `json:"inheritRotation,omitempty"`
}

// Slot 槽定义
type Slot struct {
	Name       string `json:"name"`
	Bone       string `json:"bone"`
	Color      string `json:"color,omitempty"`
	Attachment string `json:"attachment,omitempty"`
	Blend      string `json:"blend,omitempty"`
}

// RegionAttachment 区域附件
type RegionAttachment struct {
	Type      string    `json:"type,omitempty"`
	Name      string    `json:"name,omitempty"`
	Path      string    `json:"path,omitempty"`
	X         float64   `json:"x,omitempty"`
	Y         float64   `json:"y,omitempty"`
	Rotation  float64   `json:"rotation,omitempty"`
	ScaleX    float64   `json:"scaleX,omitempty"`
	ScaleY    float64   `json:"scaleY,omitempty"`
	Width     float64   `json:"width,omitempty"`
	Height    float64   `json:"height,omitempty"`
	Color     string    `json:"color,omitempty"`
	UVs       []float64 `json:"uvs,omitempty"`
	Triangles []int     `json:"triangles,omitempty"`
	Vertices  []float64 `json:"vertices,omitempty"`
}

// MeshAttachment 网格附件
type MeshAttachment struct {
	Type      string    `json:"type,omitempty"`
	Path      string    `json:"path,omitempty"`
	UVs       []float64 `json:"uvs,omitempty"`
	Triangles []int     `json:"triangles,omitempty"`
	Vertices  []float64 `json:"vertices,omitempty"`
	Edges     []int     `json:"edges,omitempty"`
	Hull      int       `json:"hull,omitempty"`
	Color     string    `json:"color,omitempty"`
	Width     float64   `json:"width,omitempty"`
	Height    float64   `json:"height,omitempty"`
}

// BoundingBoxAttachment 包围盒附件
type BoundingBoxAttachment struct {
	Type     string    `json:"type,omitempty"`
	Vertices []float64 `json:"vertices,omitempty"`
	Color    string    `json:"color,omitempty"`
}

// PathAttachment 路径附件
type PathAttachment struct {
	Type          string    `json:"type,omitempty"`
	Vertices      []float64 `json:"vertices,omitempty"`
	Lengths       []float64 `json:"lengths,omitempty"`
	Closed        bool      `json:"closed,omitempty"`
	ConstantSpeed bool      `json:"constantSpeed,omitempty"`
	Color         string    `json:"color,omitempty"`
}

// PointAttachment 点附件
type PointAttachment struct {
	Type     string  `json:"type,omitempty"`
	X        float64 `json:"x,omitempty"`
	Y        float64 `json:"y,omitempty"`
	Rotation float64 `json:"rotation,omitempty"`
	Color    string  `json:"color,omitempty"`
}

// ClippingAttachment 裁剪附件
type ClippingAttachment struct {
	Type     string    `json:"type,omitempty"`
	End      string    `json:"end,omitempty"`
	Vertices []float64 `json:"vertices,omitempty"`
	Color    string    `json:"color,omitempty"`
}

// SkinAttachments 皮肤下的附件集合
type SkinAttachments map[string]any

// Skin 皮肤定义
type Skin struct {
	Attachments map[string]SkinAttachments `json:"-"`
}

// Event 事件定义
type Event struct {
	Int    int     `json:"int,omitempty"`
	Float  float64 `json:"float,omitempty"`
	String string  `json:"string,omitempty"`
	Time   float64 `json:"time,omitempty"`
}

// IKConstraint IK约束
type IKConstraint struct {
	Name          string   `json:"name"`
	Order         int      `json:"order,omitempty"`
	Bones         []string `json:"bones"`
	Target        string   `json:"target"`
	Mix           float64  `json:"mix,omitempty"`
	Softness      float64  `json:"softness,omitempty"`
	BendDirection int      `json:"bendDirection,omitempty"`
	Compress      bool     `json:"compress,omitempty"`
	Stretch       bool     `json:"stretch,omitempty"`
	Uniform       bool     `json:"uniform,omitempty"`
}

// TransformConstraint 变换约束
type TransformConstraint struct {
	Name           string   `json:"name"`
	Order          int      `json:"order,omitempty"`
	Bones          []string `json:"bones"`
	Target         string   `json:"target"`
	TranslationX   float64  `json:"translationX,omitempty"`
	TranslationY   float64  `json:"translationY,omitempty"`
	Rotation       float64  `json:"rotation,omitempty"`
	ScaleX         float64  `json:"scaleX,omitempty"`
	ScaleY         float64  `json:"scaleY,omitempty"`
	ShearY         float64  `json:"shearY,omitempty"`
	RotateMix      float64  `json:"rotateMix,omitempty"`
	TranslateMix   float64  `json:"translateMix,omitempty"`
	ScaleMix       float64  `json:"scaleMix,omitempty"`
	ShearMix       float64  `json:"shearMix,omitempty"`
	Local          bool     `json:"local,omitempty"`
	Relative       bool     `json:"relative,omitempty"`
	OffsetRotation float64  `json:"offsetRotation,omitempty"`
	OffsetX        float64  `json:"offsetX,omitempty"`
	OffsetY        float64  `json:"offsetY,omitempty"`
	OffsetScaleX   float64  `json:"offsetScaleX,omitempty"`
	OffsetScaleY   float64  `json:"offsetScaleY,omitempty"`
	OffsetShearY   float64  `json:"offsetShearY,omitempty"`
}

// PathConstraint 路径约束
type PathConstraint struct {
	Name         string   `json:"name"`
	Order        int      `json:"order,omitempty"`
	Bones        []string `json:"bones"`
	Target       string   `json:"target"`
	Mode         string   `json:"mode,omitempty"`
	Position     float64  `json:"position,omitempty"`
	Spacing      float64  `json:"spacing,omitempty"`
	RotateMix    float64  `json:"rotateMix,omitempty"`
	TranslateMix float64  `json:"translateMix,omitempty"`
	Mix          float64  `json:"mix,omitempty"`
}

// Curve 动画曲线，可以是 "stepped"/"linear" 或贝塞尔控制点 [cx1,cy1,cx2,cy2]
type Curve struct {
	Stepped bool
	Linear  bool
	Cx1     float64
	Cy1     float64
	Cx2     float64
	Cy2     float64
}

// RotateKeyframe 旋转关键帧
type RotateKeyframe struct {
	Time  float64 `json:"time"`
	Angle float64 `json:"angle"`
	Curve any     `json:"curve,omitempty"`
}

// TranslateKeyframe 位移关键帧
type TranslateKeyframe struct {
	Time  float64 `json:"time"`
	X     float64 `json:"x,omitempty"`
	Y     float64 `json:"y,omitempty"`
	Curve any     `json:"curve,omitempty"`
}

// ScaleKeyframe 缩放关键帧
type ScaleKeyframe struct {
	Time  float64 `json:"time"`
	X     float64 `json:"x,omitempty"`
	Y     float64 `json:"y,omitempty"`
	Curve any     `json:"curve,omitempty"`
}

// ShearKeyframe 剪切关键帧
type ShearKeyframe struct {
	Time  float64 `json:"time"`
	X     float64 `json:"x,omitempty"`
	Y     float64 `json:"y,omitempty"`
	Curve any     `json:"curve,omitempty"`
}

// AttachmentKeyframe 附件切换关键帧
type AttachmentKeyframe struct {
	Time float64 `json:"time"`
	Name string  `json:"name,omitempty"`
}

// ColorKeyframe 颜色关键帧
type ColorKeyframe struct {
	Time  float64 `json:"time"`
	Color string  `json:"color,omitempty"`
	Curve any     `json:"curve,omitempty"`
}

// DeformKeyframe 变形关键帧
type DeformKeyframe struct {
	Time     float64   `json:"time"`
	Vertices []float64 `json:"vertices,omitempty"`
	Curve    any       `json:"curve,omitempty"`
	Offset   int       `json:"offset,omitempty"`
}

// EventKeyframe 事件关键帧
type EventKeyframe struct {
	Time  float64 `json:"time"`
	Name  string  `json:"name,omitempty"`
	Event *Event  `json:"event,omitempty"`
}

// DrawOrderKeyframe 绘制顺序关键帧
type DrawOrderKeyframe struct {
	Time    float64 `json:"time"`
	Offsets []struct {
		Slot   int `json:"slot"`
		Offset int `json:"offset"`
	} `json:"offsets,omitempty"`
}

// IKKeyframe IK约束关键帧
type IKKeyframe struct {
	Time         float64 `json:"time"`
	Mix          float64 `json:"mix,omitempty"`
	Softness     float64 `json:"softness,omitempty"`
	BendPositive bool    `json:"bendPositive,omitempty"`
	Compress     bool    `json:"compress,omitempty"`
	Stretch      bool    `json:"stretch,omitempty"`
	Curve        any     `json:"curve,omitempty"`
}

// TransformKeyframe 变换约束关键帧
type TransformKeyframe struct {
	Time         float64 `json:"time"`
	Rotation     float64 `json:"rotation,omitempty"`
	X            float64 `json:"x,omitempty"`
	Y            float64 `json:"y,omitempty"`
	ScaleX       float64 `json:"scaleX,omitempty"`
	ScaleY       float64 `json:"scaleY,omitempty"`
	ShearY       float64 `json:"shearY,omitempty"`
	RotateMix    float64 `json:"rotateMix,omitempty"`
	TranslateMix float64 `json:"translateMix,omitempty"`
	ScaleMix     float64 `json:"scaleMix,omitempty"`
	ShearMix     float64 `json:"shearMix,omitempty"`
	Curve        any     `json:"curve,omitempty"`
}

// PathKeyframe 路径约束关键帧
type PathKeyframe struct {
	Time         float64 `json:"time"`
	Position     float64 `json:"position,omitempty"`
	Spacing      float64 `json:"spacing,omitempty"`
	RotateMix    float64 `json:"rotateMix,omitempty"`
	TranslateMix float64 `json:"translateMix,omitempty"`
	Curve        any     `json:"curve,omitempty"`
}

// BoneAnimation 骨骼动画
type BoneAnimation struct {
	Rotate    []RotateKeyframe    `json:"rotate,omitempty"`
	Translate []TranslateKeyframe `json:"translate,omitempty"`
	Scale     []ScaleKeyframe     `json:"scale,omitempty"`
	Shear     []ShearKeyframe     `json:"shear,omitempty"`
}

// SlotAnimation 槽动画
type SlotAnimation struct {
	Attachment []AttachmentKeyframe `json:"attachment,omitempty"`
	Color      []ColorKeyframe      `json:"color,omitempty"`
}

// DeformAnimation 变形动画
type DeformAnimation struct {
	Keyframes []DeformKeyframe `json:"-"`
}

// Animation 完整动画定义
type Animation struct {
	Bones     map[string]BoneAnimation              `json:"bones,omitempty"`
	Slots     map[string]SlotAnimation              `json:"slots,omitempty"`
	DrawOrder []DrawOrderKeyframe                   `json:"drawOrder,omitempty"`
	Events    []EventKeyframe                       `json:"events,omitempty"`
	IK        map[string][]IKKeyframe               `json:"ik,omitempty"`
	Transform map[string][]TransformKeyframe        `json:"transform,omitempty"`
	Path      map[string][]PathKeyframe             `json:"path,omitempty"`
	Deform    map[string]map[string]DeformAnimation `json:"deform,omitempty"`
	FFD       map[string]map[string]DeformAnimation `json:"ffd,omitempty"`
}

// SpineData 完整的Spine骨骼动画数据
type SpineData struct {
	Skeleton   SkeletonMeta                         `json:"skeleton,omitempty"`
	Bones      []Bone                               `json:"bones,omitempty"`
	Slots      []Slot                               `json:"slots,omitempty"`
	Skins      map[string]map[string]map[string]any `json:"skins,omitempty"`
	Animations map[string]Animation                 `json:"animations,omitempty"`
	Events     map[string]Event                     `json:"events,omitempty"`
	IK         []IKConstraint                       `json:"ik,omitempty"`
	Transform  []TransformConstraint                `json:"transform,omitempty"`
	Path       []PathConstraint                     `json:"path,omitempty"`
	DrawOrder  []DrawOrderKeyframe                  `json:"drawOrder,omitempty"`
}

// ReadFile 读取并解析Spine JSON骨骼动画文件
// fileName: .json 文件路径
// 返回 SpineData，包含完整的骨骼、槽、皮肤、动画等数据
func ReadFile(fileName string) (*SpineData, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return Read(data)
}

// Read 解析Spine JSON骨骼动画数据
// data: JSON格式的骨骼动画数据
// 返回 SpineData，包含完整的骨骼、槽、皮肤、动画等数据
func Read(data []byte) (*SpineData, error) {
	var sd SpineData
	if err := json.Unmarshal(data, &sd); err != nil {
		return nil, err
	}
	return &sd, nil
}
