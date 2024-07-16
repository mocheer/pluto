package ds

import (
	"errors"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// CreateDirFromFilename 创建不存在的文件夹
func CreateDirFromFilename(fileName string) error {
	d := filepath.Dir(fileName)
	if !IsExist(d) {
		err := os.MkdirAll(d, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetFirstFiles
func GetFirstFiles(dir string) ([]fs.DirEntry, error) {
	return os.ReadDir(dir)
}

// GetFiles 获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。(不包含目录本身)
func GetFiles(dir string) (files []string, err error) {
	files = make([]string, 0, 10)
	err = filepath.Walk(dir, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if err != nil { //忽略错误
			return err
		}
		if !fi.IsDir() { // 忽略目录
			files = append(files, filename)
		}

		return nil
	})
	return files, err
}

// GetDirs
func GetDirs(dir string) (dirs []string, err error) {
	dirs = make([]string, 0, 10)
	err = filepath.Walk(dir, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if err != nil { //忽略错误
			return err
		}
		if fi.IsDir() { // 忽略目录
			dirs = append(dirs, filename)
		}
		return nil
	})
	return dirs, err
}

func GetFilesWithSuffix(dir, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
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

// CopyDir 拷贝文件夹和所有子文件夹
func CopyDir(source string, dst string) error {
	srcinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// 创建目标目录
	err = os.MkdirAll(dst, srcinfo.Mode())
	if err != nil {
		return err
	}

	dir, _ := os.Open(source)
	obs, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

	var errs []error
	for _, obj := range obs {
		fsource := source + "/" + obj.Name()
		fdest := dst + "/" + obj.Name()
		if obj.IsDir() {
			// 循环拷贝子文件夹
			err = CopyDir(fsource, fdest)
			if err != nil {
				errs = append(errs, err)
			}
		} else {
			// 拷贝文件
			err = CopyFile(fsource, fdest)
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	if len(errs) > 0 {
		var errString string
		for _, err := range errs {
			errString += err.Error() + "\n"
		}
		return errors.New(errString)
	}

	return nil
}

// RemoveEmptyDir 移除空目录
func RemoveEmptyDir(dir string) (bool, error) {
	de, err := os.ReadDir(dir)
	if err != nil {
		return false, err
	}
	if len(de) > 0 {
		for _, info := range de {
			if info.IsDir() {
				isDirEmpty, err := RemoveEmptyDir(path.Join(dir, info.Name()))
				if err != nil {
					return false, err
				}
				if !isDirEmpty {
					return false, nil
				}
			} else {
				return false, nil
			}
		}
	}
	os.RemoveAll(dir)
	return true, nil
}
