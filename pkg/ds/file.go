package ds

import (
	"io"
	"os"
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
	err := CreateDirFromFilename(fileName)
	if err != nil {
		return nil, err
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
// 慎用，如果文件存在，该方法不会创建新文件，所以如果直接写入data数据时，当文件内容大于data的数据长度时，会保留后面的数据
func OpenOrCreate(fileName string) (*os.File, error) {
	err := CreateDirFromFilename(fileName) //确保目录存在，不存在的话会报错
	if err != nil {
		return nil, err
	}
	// O_RDWR：可读可写
	// O_CREATE：如果不存在将创建一个新文件
	return os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
}

// Save 保存
func Save(fileName string, data []byte) error {
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
