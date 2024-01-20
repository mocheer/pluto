package ds

import "path/filepath"

func GetFilenameNotExt(filename string) string {
	baseName := filepath.Base(filename) // 获取基本名称
	extension := filepath.Ext(filename) // 获取后缀
	filename = baseName[0 : len(baseName)-len(extension)]
	return filename
}
