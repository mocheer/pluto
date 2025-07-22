package img_test

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/mocheer/pluto/pkg/ts/img"
)

// TestSave
func TestSave(t *testing.T) {
	p, _, err := img.FromFile("test.png")
	if err != nil {
		t.Error(err)
	}
	p.Save("test_save.png", "png")
}

func TestBase64(t *testing.T) {
	p, f, err := img.FromFile("testdata/auth.jpg")
	if err != nil {
		t.Error(err)
	}
	t.Log(f)
	t.Log(p.ToBase64())
}

// TestColor
func TestColor(t *testing.T) {
	c := color.RGBA{255, 16, 16, 254}
	r, g, b, a := c.RGBA()
	t.Log(r, g, b, a)
	var hexr, hexg, hexb string
	hexr = strconv.FormatUint(uint64(r), 16)
	if r < 16 {
		hexr = "0" + hexr
	}
	hexg = strconv.FormatUint(uint64(g), 16)
	if g < 16 {
		hexg = "0" + hexg
	}
	hexb = strconv.FormatUint(uint64(b), 16)
	if b < 16 {
		hexb = "0" + hexb
	}

	t.Log(hexr + hexg + hexb)
}

func TestJpeg2Png(t *testing.T) {
	img.ConvertJPEGToTransparentPNG("./testdata/heatmap.jpg", "./testdata/heatmap.png", color.RGBA{194, 193, 191, 255}, 10)
}

func TestEdge(t *testing.T) {
	// 解析命令行参数
	inputPtr := "./testdata/fj.png"
	outputPtr := "./testdata/fj.edge.jpeg"

	// 打开并解码图片
	file, err := os.Open(inputPtr)
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	defer file.Close()

	img2, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("解码图片失败: %v", err)
	}

	saveJpg(img.Edge(img2), outputPtr)
}

// savePng 保存图像到文件
func savePng(img image.Image, outputPath string) {
	// 创建输出目录（如果不存在）
	dir := filepath.Dir(outputPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("无法创建输出文件: %v", err)
	}
	defer file.Close()

	// 保存为PNG
	if err := png.Encode(file, img); err != nil {
		log.Fatalf("保存图片失败: %v", err)
	}

}

// savePng 保存图像到文件
func saveJpg(img image.Image, outputPath string) {
	// 创建输出目录（如果不存在）
	dir := filepath.Dir(outputPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("无法创建输出文件: %v", err)
	}
	defer file.Close()

	// 保存为
	if err := jpeg.Encode(file, img, &jpeg.Options{Quality: 75}); err != nil {
		log.Fatalf("保存图片失败: %v", err)
	}

}
