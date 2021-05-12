package fs

import (
	"image"
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

// Create
func Create(name string) (*os.File, error) {
	err := os.MkdirAll(filepath.Dir(name), os.ModePerm)
	if err == nil {
		file, err := os.Create(name)
		return file, err
	}
	return nil, err
}

// OpenOrCreate 创建不存在的文件
func OpenOrCreate(name string, flag int, perm os.FileMode) (*os.File, error) {
	isExit := IsExist(name)
	if !isExit {
		return Create(name)
	}
	return os.OpenFile(name, flag, perm)
}

// MkdirNotExist 创建不存在的文件夹
func MkdirNotExist(path string) error {
	isExit := IsExist(path)
	if !isExit {
		err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// EachDir 获取指定目录及所有子目录下的所有文件
func EachDir(dir string, fn func(filename string)) error {
	err := filepath.Walk(dir, func(filename string, fi os.FileInfo, err error) error {
		if err != nil { //忽略错误
			return err
		}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		fn(filename)
		return nil
	})
	return err
}

// EachDirToRemove 遍历文件夹，当回调函数返回true的时候删除文件
// 注意，当一个目录下所有文件都被删除，这个目录本身也不会被删除
func EachDirToRemove(dir string, fn func(filename string) bool) error {
	return EachDir(dir, func(filename string) {
		if fn(filename) {
			os.Remove(filename)
		}
	})
}

// WalkDir 获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDir(dir, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)                                                  //忽略后缀匹配的大小写
	err = filepath.Walk(dir, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if err != nil { //忽略错误
			return err
		}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

// Append 往文件尾部添加字符串
func Append(path string, content string) {
	f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	f.Write([]byte(content))
	f.Close()
}

// EachDirAppendHead 遍历目录下的所有文件，添加文件头
func EachDirAppendHead(dir string, content string, options map[string]interface{}) error {
	return EachDir(dir, func(filename string) {
		if options != nil {
			if options["suffix"] != nil {
				if !strings.HasSuffix(filename, options["suffix"].(string)) {
					return
				}
			}
		}
		AppendHead(filename, content)
	})
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

// GetImageFromPath 读取图片返回image对象
func GetImageFromPath(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	image, _, err := image.Decode(f)
	return image, err
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
