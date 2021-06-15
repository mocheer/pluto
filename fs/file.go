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
func MustReadFile(fileName string) []byte {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return data
}

// Create
func Create(fileName string) (*os.File, error) {
	err := os.MkdirAll(filepath.Dir(fileName), os.ModePerm)
	if err == nil {
		file, err := os.Create(fileName)
		return file, err
	}
	return nil, err
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

// Append 往文件尾部添加字符串
func Append(fileName string, content string) {
	f, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	f.Write([]byte(content))
	f.Close()
}

// Append 往文件尾部添加字符串
func AppendHead(fileName string, content string) {
	f, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	defer f.Close()
	data, err := io.ReadAll(f)
	old := string(data)
	if err == nil && !strings.Contains(old, content) {
		f.WriteAt([]byte(content+"\n"+old), 0)
	}
}

// SaveFile 保存图片
func SaveFile(fileName string, data []byte) error {
	f, err := OpenOrCreate(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err == nil {
		f.Write(data)
	}
	defer f.Close()
	return err
}
