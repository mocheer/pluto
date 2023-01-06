package ds

import (
	"io"
	"os"
	"path/filepath"
)

func ReadFile(fileName string) ([]byte, error) {
	return os.ReadFile(fileName)
}

func MustReadFile(fileName string) []byte {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return data
}

// Create 创建一个不存在的文件(已存在则忽略)
func Create(fileName string) (*os.File, error) {
	dir := filepath.Dir(fileName)
	if !IsExist(dir) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	file, err := os.Create(fileName)
	return file, err
}

// MustCreate 创建文件
func MustCreate(fileName string) *os.File {
	f, err := Create(fileName)
	if err != nil {
		panic(err)
	}
	return f
}

// OpenOrCreate
func OpenOrCreate(fileName string) (*os.File, error) {
	// O_RDWR：可读可写
	// O_CREATE：如果不存在将创建一个新文件
	return os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
}

// Save 保存
func Save(fileName string, data []byte) error {
	//这里不用 OpenOrCreate，因为有问题的，当文件有内容且内容大于data长度时，会保留
	f, err := Create(fileName)
	if err == nil {
		f.Write(data)
	}
	defer f.Close()
	return err
}

// CopyFile 拷贝文件
func CopyFile(src, dst string) (err error) {
	sf, err := os.Open(src)
	if err != nil {
		return
	}
	defer sf.Close()
	//
	df, err := Create(dst)
	if err != nil {
		return
	}
	defer df.Close()
	_, err = io.Copy(df, sf)
	return
}
