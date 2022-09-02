package ds_gzip

import (
	"compress/gzip"
	"io"

	"os"
)

// Read 读取文件并生成对应的gzip字节数组
func Read(filename string) ([]byte, error) {
	// 打开本地gz格式压缩包
	fr, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	// defer: 在函数退出时,执行关闭文件
	defer fr.Close()
	// 创建gzip文件读取对象
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return nil, err
	}
	// defer: 在函数退出时,执行关闭gzip对象
	defer gr.Close()
	// 读取gzip对象内容
	rBuf, err := io.ReadAll(gr)
	if err != nil {
		return nil, err
	}
	return rBuf, nil
}
