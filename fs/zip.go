package fs

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path"
	"strings"
)

// EachZip 遍历zip文件
func EachZip(fileName string, callback func(*zip.File)) {
	// 读取
	zipFile, err := zip.OpenReader(fileName)
	if err != nil {
		panic(err.Error())
	}
	defer zipFile.Close()
	// 遍历所有文件
	for _, f := range zipFile.File {
		callback(f)
	}
}

// ExtractZip 解压缩zip文件
func ExtractZip(fileName string) {
	// 读取
	EachZip(fileName, func(f *zip.File) {
		info := f.FileInfo()
		if info.IsDir() {
			err := os.MkdirAll(f.Name, os.ModePerm)
			if err != nil {
				panic(err.Error())
			}
			return
		}
		srcFile, err := f.Open()
		if err != nil {
			panic(err.Error())
		}
		defer srcFile.Close()

		newFile, err := os.Create(f.Name)
		if err != nil {
			panic(err.Error())
		}
		defer newFile.Close()

		io.Copy(newFile, srcFile)
	})

}

// WriteZip 读取某个文件并创建压缩包
func WriteZip(data []byte, fileName string) {
	// 缓存压缩文件内容
	buf := new(bytes.Buffer)

	// 创建zip
	writer := zip.NewWriter(buf)
	defer writer.Close()

	// 接收
	f, _ := writer.Create(fileName)
	f.Write(data)

	fileName = strings.TrimSuffix(fileName, path.Ext(fileName)) + ".zip"
	os.WriteFile(fileName, buf.Bytes(), 0644)
}
