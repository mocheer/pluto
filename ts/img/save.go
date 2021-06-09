package img

import (
	"encoding/base64"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

// Save 保存为文件
func (p *Picture) Save(path string) error {
	switch p.Type {
	case JPEG:
		return SaveAsJPEG(p.Image, path)
	case PNG:
		return SaveAsPNG(p.Image, path)
	case GIF:
		return SaveAsGIF(p.Image, path)
	}
	return errors.New("不支持的图片格式")
}

// SaveAsJPEG
func SaveAsJPEG(target image.Image, path string) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}
	defer file.Close()
	return jpeg.Encode(file, target, nil)
}

// SaveAsPNG
func SaveAsPNG(target image.Image, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, target)
}

// SaveAsGIF
func SaveAsGIF(target image.Image, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return gif.Encode(file, target, nil)
}

// SaveBase64 将base64字串保存为图片
func SaveBase64(source string, path string) (err error) {
	data, err := base64.StdEncoding.DecodeString(source) //图片文件并把文件写入到buffer
	if err != nil {
		return
	}
	err = os.WriteFile(path, data, 0666) //buffer输出到png文件中,0666 访问权限
	return
}
