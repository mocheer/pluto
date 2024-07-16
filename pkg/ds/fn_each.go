package ds

import (
	"os"
	"path/filepath"
	"strings"
)

// Each 遍历指定目录及所有子目录下的所有文件(包括目录本身)
func Each(dir string, fn func(filename string, fi os.FileInfo)) error {
	err := filepath.Walk(dir, func(filename string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fn(filename, fi)
		return nil
	})
	return err
}

// EachFiles 遍历指定目录及所有子目录下的所有文件
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

// EachFilesToRemove 遍历文件夹，当回调函数返回true的时候删除文件
// 注意，当一个目录下所有文件都被删除，这个目录本身也不会被删除
func EachFilesToRemove(dir string, fn func(filename string, fi os.FileInfo) bool) error {
	return EachFiles(dir, func(filename string, fi os.FileInfo) {
		if fn(filename, fi) {
			os.Remove(filename)
		}
	})
}

// 批量重命名指定目录及所有子目录下的所有文件。(包含目录)
// func EachToRename(dir string, fn func(oldName string) string) (err error) {
// 	return Each(dir, func(filename string, fi os.FileInfo) {
// 		oldName := fi.Name()
// 		newName := fn(oldName)
// 		os.Rename(filename, filepath.Join(filepath.Dir(oldName), newName))
// 	})
// }

// 批量重命名指定目录及所有子目录下的所有目录
func EachDirsToRename(dir string, fn func(oldName string) string) (err error) {
	return Each(dir, func(filename string, fi os.FileInfo) {
		if fi.IsDir() { // 忽略目录
			oldName := fi.Name()
			newName := fn(oldName)
			os.Rename(filename, filepath.Join(filepath.Dir(filename), newName))
		}
	})
}

// 批量重命名指定目录及所有子目录下的所有文件。(不包含目录)
func EachFilesToRename(dir string, fn func(oldName string) string) (err error) {
	return EachFiles(dir, func(filename string, fi os.FileInfo) {
		oldName := fi.Name()
		newName := fn(oldName)
		os.Rename(filename, filepath.Join(filepath.Dir(filename), newName))
	})
}

// EachFilesToAppendHead 遍历目录下的所有文件，添加文件头
func EachFilesToAppendHead(dir string, content string, options map[string]any) error {
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
