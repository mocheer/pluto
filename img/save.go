package img

import (
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

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

// SaveBase64 将base64字串保存为图片
func SaveBase64(source string, filePath string) (err error) {
	data, err := base64.StdEncoding.DecodeString(source) //图片文件并把文件写入到buffer
	if err != nil {
		return
	}
	err = os.WriteFile(filePath, data, 0666) //buffer输出到png文件中,0666 访问权限
	return
}
