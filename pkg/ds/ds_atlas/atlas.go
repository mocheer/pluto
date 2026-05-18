// Package atlas 提供 TexturePacker 导出的 .atlas 合图拆解功能。
//
// 功能概述：
//  1. 读取 .atlas JSON 元数据文件
//  2. 根据元数据中的 image 字段加载对应的大图
//  3. 遍历 frames，按帧信息从大图中裁剪出子图
//  4. 返回 Atlas 对象，包含所有大图和帧数据
//
// .atlas 文件格式说明（TexturePacker 通用格式）：
//
//	{
//	  "meta": { "image": "sheet.png" },         // 大图文件名（多张用逗号分隔）
//	  "frames": {
//	    "hero.png": {                            // 帧名称
//	      "frame": { "idx": 0, "x": 10, "y": 20, "w": 64, "h": 64 }
//	    }
//	  }
//	}
package atlas

import (
	"encoding/json"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"strings"
)

// parse 处理单个 atlas 文件，返回 Atlas 对象
func parse(altasName string) (*Atlas, error) {
	content, err := os.ReadFile(altasName)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", altasName, err)
	}

	var data atlasConfig
	if err := json.Unmarshal(content, &data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON %s: %w", altasName, err)
	}

	var bigImages []image.Image
	// 当不存在meta.Image时，默认使用与.atlas文件名相同的文件名
	if data.Meta.Image == "" {
		data.Meta.Image = strings.ReplaceAll(filepath.Base(altasName), ".altas", ".png")
	}
	dirName := filepath.Dir(altasName)
	imagePs := strings.Split(data.Meta.Image, ",")
	for _, ipath := range imagePs {
		imPath := filepath.Join(dirName, ipath)
		// read image file
		f, err := os.OpenFile(imPath, os.O_RDONLY, 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to open image %s: %w", imPath, err)
		}
		defer f.Close()
		im, t, err := image.Decode(f)
		if err != nil {
			return nil, fmt.Errorf("failed to decode image %s as %s: %w", imPath, t, err)
		}
		bigImages = append(bigImages, im)
	}

	frames := make(map[string]Frame, len(data.Frames))
	for k, frameInfo := range data.Frames {
		fr := frameInfo.Frame
		im := bigImages[fr.Idx]
		sub := im.(subImager).SubImage(image.Rect(fr.X, fr.Y, fr.X+fr.W, fr.Y+fr.H))

		frames[k] = Frame{
			Frame:    fr,
			SubImage: sub,
		}
	}
	return &Atlas{
		Images: bigImages,
		Frames: frames,
	}, nil
}
