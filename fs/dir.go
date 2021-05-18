package fs

import (
	"os"
	"path/filepath"
	"strings"
)

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

// EachFiles 获取指定目录及所有子目录下的所有文件
func EachFiles(dir string, fn func(filename string, fi os.FileInfo)) error {
	err := filepath.Walk(dir, func(filename string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		fn(filename, fi)
		return nil
	})
	return err
}

// EachFilesRemove 遍历文件夹，当回调函数返回true的时候删除文件
// 注意，当一个目录下所有文件都被删除，这个目录本身也不会被删除
func EachFilesRemove(dir string, fn func(filename string, fi os.FileInfo) bool) error {
	return EachFiles(dir, func(filename string, fi os.FileInfo) {
		if fn(filename, fi) {
			os.Remove(filename)
		}
	})
}

// 批量重命名指定目录及所有子目录下的所有文件。(不会重命名目录文件)
func EachFilesRename(dir string, fn func(oldName string) string) (err error) {
	return EachFiles(dir, func(filename string, fi os.FileInfo) {
		oldName := fi.Name()
		newName := fn(oldName)
		os.Rename(filename, filepath.Join(filepath.Dir(oldName), newName))
	})
}

// EachFilesAppendHead 遍历目录下的所有文件，添加文件头
func EachFilesAppendHead(dir string, content string, options map[string]interface{}) error {
	return EachFiles(dir, func(filename string, fi os.FileInfo) {
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

// GetFiles 获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func GetFiles(dir, suffix string) (files []string, err error) {
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
