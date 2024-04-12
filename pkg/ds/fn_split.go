package ds

import (
	"fmt"
	"io"
	"os"
)

// 百度云盘、阿里云盘不支持大文件上传，这里用一个简单的分割程序切分成小文件

// SplitFile 将任意格式的大文件分割成固定大小的小文件
func SplitFile(inputFile string, chunkSize int) error {
	// 打开要分割的大文件进行读取
	file, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// 设置用于保存分割文件的目录
	outputDir := fmt.Sprintf("./%s_chunks/", file.Name())
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}

	chunkNum := 0
	for {
		// 创建分割文件
		chunkPath := fmt.Sprintf("%s%d", outputDir, chunkNum)
		outputFile, err := os.Create(chunkPath)
		if err != nil {
			return err
		}

		// 流式复制，每次都是从上一次复制的位置继续往下读取
		// 使用 io.CopyN 限制读取的大小为 chunkSize
		_, err = io.CopyN(outputFile, file, int64(chunkSize))
		if err != nil {
			outputFile.Close()
			if err == io.EOF {
				break
			}
			return err
		}

		outputFile.Close()
		chunkNum++
	}

	return nil
}

// MergeFiles 将分割后的文件合并成原始大文件
// 需要注意chunkPaths合并的文件顺序
func MergeFiles(chunkPaths []string, outputFile string) error {
	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	for _, chunkPath := range chunkPaths {
		inFile, err := os.Open(chunkPath)
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, inFile)
		inFile.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
