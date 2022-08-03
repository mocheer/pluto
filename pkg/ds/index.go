package ds

import (
	"os"
	"path"
)

// IsExist 检查文件或目录是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// Copy 拷贝文件或目录
func Copy(src, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return CopyDir(src, dst)
	}

	return CopyFile(src, dst)
}

// Clear 清空目录，包括目录本身
func Clear(dst string) error {
	return os.RemoveAll(dst)
}

// ClearDir 清空目录，不包括目录本身
func ClearDir(dst string) error {
	dir, err := os.ReadDir(dst)
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{dst, d.Name()}...))
	}
	return err
}

func Remove(fileName string) error {
	return os.Remove(fileName)
}
