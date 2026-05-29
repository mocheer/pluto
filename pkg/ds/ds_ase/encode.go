package ds_ase

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"image"
	"image/draw"
)

// Encode 编码aseprite文件
// 将 Aseprite JSON 数据 + 精灵表图像编码为 .aseprite v1.3 字节流
// TODO 测试编码是否正确
func Encode(data *Aseprite, sheet image.Image) ([]byte, error) {
	// 1. 确定统一画布大小：取所有帧 sourceSize 的最大值
	canvasW, canvasH := 0, 0
	for _, f := range data.Frames {
		w := f.SourceSize.W
		h := f.SourceSize.H
		if f.Trimmed {
			// 裁剪帧的源尺寸已包含完整画布
		}
		if w > canvasW {
			canvasW = w
		}
		if h > canvasH {
			canvasH = h
		}
	}
	if canvasW == 0 || canvasH == 0 {
		return nil, fmt.Errorf("invalid canvas size")
	}

	// 2. 构建图层（这里只有一个图层，可根据需求扩展）
	layers := []layer{
		{
			flags:      1, // visible
			typ:        layerTypeNormal,
			childLevel: 0,
			blendMode:  blendModeNormal,
			opacity:    255,
			name:       "Layer 1",
		},
	}

	// 3. 构建标签（帧标签）
	var tags []frameTag
	for _, t := range data.Meta.FrameTags {
		dir := uint8(0)
		switch t.Direction {
		case "reverse":
			dir = 1
		case "pingpong":
			dir = 2
		}
		tags = append(tags, frameTag{
			from:    uint16(t.From),
			to:      uint16(t.To),
			loopDir: dir,
			name:    t.Name,
		})
	}

	// 4. 准备每一帧的 Cel 数据

	var frames []frameData
	for _, fi := range data.Frames {
		// 4.1 从精灵表裁剪子图
		srcRect := image.Rect(fi.Frame.X, fi.Frame.Y, fi.Frame.X+fi.Frame.W, fi.Frame.Y+fi.Frame.H)
		subImg := image.NewRGBA(srcRect)
		draw.Draw(subImg, srcRect, sheet, srcRect.Min, draw.Src)

		// 4.2 处理旋转（逆时针 90 度，恢复为原始方向）
		var oriented *image.RGBA
		if fi.Rotated {
			oriented = rotateCounterClockwise90(subImg)
		} else {
			oriented = subImg
		}

		// 4.3 处理裁剪：将子图放置在源尺寸画布的指定偏移处
		var finalPixels []byte
		var finalW, finalH uint16
		var offsetX, offsetY int16

		if fi.Trimmed {
			// 创建尺寸为 SourceSize 的画布
			canvas := image.NewRGBA(image.Rect(0, 0, fi.SourceSize.W, fi.SourceSize.H))
			draw.Draw(canvas,
				image.Rect(fi.SpriteSourceSize.X, fi.SpriteSourceSize.Y,
					fi.SpriteSourceSize.X+fi.SpriteSourceSize.W,
					fi.SpriteSourceSize.Y+fi.SpriteSourceSize.H),
				oriented, image.Point{}, draw.Src)
			finalW = uint16(fi.SourceSize.W)
			finalH = uint16(fi.SourceSize.H)
			offsetX = 0
			offsetY = 0
			finalPixels = rgbaPixels(canvas)
		} else {
			finalW = uint16(fi.SourceSize.W)
			finalH = uint16(fi.SourceSize.H)
			offsetX = 0
			offsetY = 0
			finalPixels = rgbaPixels(oriented)
		}

		frames = append(frames, frameData{
			duration: uint32(fi.Duration),
			cels: []cel{
				{
					layerIndex: 0,
					x:          offsetX,
					y:          offsetY,
					opacity:    255,
					celType:    celTypeCompressed,
					width:      finalW,
					height:     finalH,
					pixels:     finalPixels,
				},
			},
		})
	}

	// 5. 写入二进制数据
	return writeFile(canvasW, canvasH, layers, frames, tags)
}

type frameData struct {
	duration uint32
	cels     []cel
}

// ---------------------------------------------------------------------------
// 与 .aseprite 格式相关的常量
// ---------------------------------------------------------------------------
const (
	magicNumber      uint16 = 0xA5E0
	frameMagic       uint16 = 0xF1FA
	chunkTypeLayer   uint16 = 0x2004
	chunkTypeCel     uint16 = 0x2005
	chunkTypeTags    uint16 = 0x2017
	chunkTypePalette uint16 = 0x2018

	celTypeRaw        uint16 = 0
	celTypeCompressed uint16 = 2
	celTypeLinked     uint16 = 1

	layerTypeNormal uint16 = 0
	blendModeNormal uint16 = 0
	colorDepthRGBA  uint16 = 32

	flagLayerOpacityValid uint32 = 1
)

// ---------------------------------------------------------------------------
// 内部结构（方便组装文件）
// ---------------------------------------------------------------------------
type layer struct {
	flags      uint16
	typ        uint16
	childLevel uint16
	blendMode  uint16
	opacity    uint8
	name       string
}

type cel struct {
	layerIndex uint16
	x, y       int16
	opacity    uint8
	celType    uint16
	width      uint16
	height     uint16
	pixels     []byte // RGBA
	frameLink  uint16 // for linked cel
}

type frameTag struct {
	from, to uint16
	loopDir  uint8
	name     string
}

// ---------------------------------------------------------------------------
// 写入函数
// ---------------------------------------------------------------------------
func writeFile(canvasW, canvasH int, layers []layer, frames []frameData, tags []frameTag) ([]byte, error) {
	out := new(bytes.Buffer)

	// 占位头
	header := make([]byte, 128)
	if _, err := out.Write(header); err != nil {
		return nil, err
	}

	// 第一帧的全局块
	var globalChunks []chunkBuilder
	for _, l := range layers {
		globalChunks = append(globalChunks, chunkBuilder{chunkTypeLayer, layerChunkData(l)})
	}
	if len(tags) > 0 {
		globalChunks = append(globalChunks, chunkBuilder{chunkTypeTags, tagsChunkData(tags)})
	}

	for i, fd := range frames {
		var frameChunks []chunkBuilder
		if i == 0 {
			frameChunks = append(frameChunks, globalChunks...)
		}
		for _, c := range fd.cels {
			celData, err := celChunkData(c)
			if err != nil {
				return nil, fmt.Errorf("frame %d cel: %w", i, err)
			}
			frameChunks = append(frameChunks, chunkBuilder{chunkTypeCel, celData})
		}
		if err := writeFrame(out, fd.duration, frameChunks); err != nil {
			return nil, err
		}
	}

	// 回填文件头
	raw := out.Bytes()
	fileSize := uint32(len(raw))
	headerData := buildHeader(fileSize, uint16(canvasW), uint16(canvasH), uint16(len(frames)))
	copy(raw[0:128], headerData)
	return raw, nil
}

func buildHeader(fileSize uint32, width, height uint16, frames uint16) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, fileSize)
	binary.Write(buf, binary.LittleEndian, magicNumber)
	binary.Write(buf, binary.LittleEndian, frames)
	binary.Write(buf, binary.LittleEndian, width)
	binary.Write(buf, binary.LittleEndian, height)
	binary.Write(buf, binary.LittleEndian, colorDepthRGBA)
	binary.Write(buf, binary.LittleEndian, flagLayerOpacityValid)
	binary.Write(buf, binary.LittleEndian, uint16(0)) // speed (deprecated)
	binary.Write(buf, binary.LittleEndian, uint32(0)) // reserved
	binary.Write(buf, binary.LittleEndian, uint32(0)) // reserved
	binary.Write(buf, binary.LittleEndian, uint8(0))  // transparent index
	buf.Write([]byte{0, 0, 0})
	binary.Write(buf, binary.LittleEndian, uint16(0)) // num colors (0 = 256)
	buf.Write([]byte{1})                              // pixel width
	buf.Write([]byte{1})                              // pixel height
	binary.Write(buf, binary.LittleEndian, int16(0))  // grid x
	binary.Write(buf, binary.LittleEndian, int16(0))  // grid y
	binary.Write(buf, binary.LittleEndian, uint16(0)) // grid width
	binary.Write(buf, binary.LittleEndian, uint16(0)) // grid height
	buf.Write(make([]byte, 84))                       // reserved
	return buf.Bytes()
}

type chunkBuilder struct {
	chunkType uint16
	data      []byte
}

func writeFrame(buf *bytes.Buffer, duration uint32, chunks []chunkBuilder) error {
	frameData := new(bytes.Buffer)
	binary.Write(frameData, binary.LittleEndian, frameMagic)
	binary.Write(frameData, binary.LittleEndian, uint16(0)) // old chunks count
	binary.Write(frameData, binary.LittleEndian, duration)
	frameData.Write([]byte{0, 0}) // reserved
	chunkCount := uint32(len(chunks))
	binary.Write(frameData, binary.LittleEndian, chunkCount)

	for _, c := range chunks {
		writeChunk(frameData, c.chunkType, c.data)
	}

	frameSize := uint32(frameData.Len())
	binary.Write(buf, binary.LittleEndian, frameSize)
	buf.Write(frameData.Bytes())
	return nil
}

func writeChunk(buf *bytes.Buffer, chunkType uint16, data []byte) {
	chunkSize := uint32(2 + len(data))
	binary.Write(buf, binary.LittleEndian, chunkSize)
	binary.Write(buf, binary.LittleEndian, chunkType)
	buf.Write(data)
}

func layerChunkData(l layer) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, l.flags)
	binary.Write(buf, binary.LittleEndian, l.typ)
	binary.Write(buf, binary.LittleEndian, l.childLevel)
	binary.Write(buf, binary.LittleEndian, uint16(0)) // default width
	binary.Write(buf, binary.LittleEndian, uint16(0)) // default height
	binary.Write(buf, binary.LittleEndian, l.blendMode)
	buf.WriteByte(l.opacity)
	buf.Write([]byte{0, 0, 0}) // reserved
	nameBytes := []byte(l.name)
	binary.Write(buf, binary.LittleEndian, uint16(len(nameBytes)))
	buf.Write(nameBytes)
	return buf.Bytes()
}

func tagsChunkData(tags []frameTag) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, uint16(len(tags)))
	buf.Write(make([]byte, 8)) // reserved
	for _, t := range tags {
		binary.Write(buf, binary.LittleEndian, t.from)
		binary.Write(buf, binary.LittleEndian, t.to)
		buf.WriteByte(t.loopDir)
		buf.Write(make([]byte, 8)) // reserved
		nameBytes := []byte(t.name)
		binary.Write(buf, binary.LittleEndian, uint16(len(nameBytes)))
		buf.Write(nameBytes)
	}
	return buf.Bytes()
}

func celChunkData(c cel) ([]byte, error) {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, c.layerIndex)
	binary.Write(buf, binary.LittleEndian, c.x)
	binary.Write(buf, binary.LittleEndian, c.y)
	buf.WriteByte(c.opacity)
	binary.Write(buf, binary.LittleEndian, c.celType)
	buf.Write(make([]byte, 7)) // reserved

	switch c.celType {
	case celTypeRaw:
		binary.Write(buf, binary.LittleEndian, c.width)
		binary.Write(buf, binary.LittleEndian, c.height)
		buf.Write(c.pixels)
	case celTypeCompressed:
		binary.Write(buf, binary.LittleEndian, c.width)
		binary.Write(buf, binary.LittleEndian, c.height)
		var compBuf bytes.Buffer
		w := zlib.NewWriter(&compBuf)
		if _, err := w.Write(c.pixels); err != nil {
			return nil, err
		}
		w.Close()
		buf.Write(compBuf.Bytes())
	case celTypeLinked:
		binary.Write(buf, binary.LittleEndian, c.frameLink)
	default:
		return nil, fmt.Errorf("unsupported cel type %d", c.celType)
	}
	return buf.Bytes(), nil
}

// ---------------------------------------------------------------------------
// 图像处理工具
// ---------------------------------------------------------------------------

// rotateCounterClockwise90 逆时针旋转 90 度
func rotateCounterClockwise90(img *image.RGBA) *image.RGBA {
	srcBounds := img.Bounds()
	w, h := srcBounds.Dx(), srcBounds.Dy()
	dst := image.NewRGBA(image.Rect(0, 0, h, w))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			newX := y
			newY := w - 1 - x
			dst.SetRGBA(newX, newY, img.RGBAAt(x+srcBounds.Min.X, y+srcBounds.Min.Y))
		}
	}
	return dst
}

// rgbaPixels 从 *image.RGBA 提取原始像素字节
func rgbaPixels(img *image.RGBA) []byte {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	pixels := make([]byte, w*h*4)
	for y := 0; y < h; y++ {
		offset := y * w * 4
		copy(pixels[offset:], img.Pix[img.PixOffset(0, y):img.PixOffset(w, y)])
	}
	return pixels
}
