package fs

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

// IsExist 检查文件或目录是否存在
func IsExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

//
func MustReadFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return data
}

// Create
func Create(name string) (*os.File, error) {
	err := os.MkdirAll(filepath.Dir(name), os.ModePerm)
	if err == nil {
		file, err := os.Create(name)
		return file, err
	}
	return nil, err
}

// MustCreate
func MustCreate(name string) *os.File {
	f, err := Create(name)
	if err != nil {
		panic(err)
	}
	return f
}

// OpenOrCreate 创建不存在的文件
func OpenOrCreate(name string, flag int, perm os.FileMode) (*os.File, error) {
	isExit := IsExist(name)
	if !isExit {
		return Create(name)
	}
	return os.OpenFile(name, flag, perm)
}

// Append 往文件尾部添加字符串
func Append(path string, content string) {
	f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	f.Write([]byte(content))
	f.Close()
}

// Append 往文件尾部添加字符串
func AppendHead(path string, content string) {
	f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	defer f.Close()
	data, err := io.ReadAll(f)
	old := string(data)
	if err == nil && !strings.Contains(old, content) {
		f.WriteAt([]byte(content+"\n"+old), 0)
	}
}

// SaveFile 保存图片
func SaveFile(path string, data []byte) error {
	f, err := OpenOrCreate(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err == nil {
		f.Write(data)
	}
	defer f.Close()
	return err
}
