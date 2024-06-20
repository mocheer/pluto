package ds_zip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Extract 解压缩zip文件到指定目录
func Extract(fileName string, target string) {
	// 读取
	Each(fileName, func(f *zip.File) {
		info := f.FileInfo()
		// 解压路径
		targetName := filepath.Join(target, f.Name)
		if info.IsDir() {
			err := os.MkdirAll(targetName, os.ModePerm)
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

		newFile, err := os.Create(targetName)
		if err != nil {
			panic(err.Error())
		}
		defer newFile.Close()

		io.Copy(newFile, srcFile)
	})

}

// ExtractCurrent 解压缩zip文件到当前目录
func ExtractCurrent(fileName string) {
	Extract(fileName, "")
}
