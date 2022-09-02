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

// OpenOrCreate 创建不存在的文件
func OpenOrCreate(fileName string, flag int, perm os.FileMode) (*os.File, error) {
	isExit := IsExist(fileName)
	if !isExit {
		return Create(fileName)
	}
	return os.OpenFile(fileName, flag, perm)
}

// Save 保存图片
func Save(fileName string, data []byte) error {
	f, err := OpenOrCreate(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
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
	df, err := OpenOrCreate(dst, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer df.Close()
	_, err = io.Copy(df, sf)
	return
}
