package ds

import (
	"os"
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
