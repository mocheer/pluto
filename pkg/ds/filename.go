package ds

import (
	"net/url"
	"path"
	"path/filepath"

	"github.com/mocheer/pluto/pkg/fn"
)

func GetFilenameNotExt(filename string) string {
	baseName := filepath.Base(filename) // 获取基本名称
	extension := filepath.Ext(filename) // 获取后缀
	filename = baseName[0 : len(baseName)-len(extension)]
	return filename
}

// SafeFileName
// 经常会有不规范的文件名，文件名不能包含的字符  \/:*?"<>|，不能是空字符串，点号开头
// 经常需要将url路径的内容下载保存，这里统一的文件名修正
// 这里会自动去掉目录，只保留文件名
// @see govalidator.SafeFileName
// 文件名不常用的字符有：~{}[]`^@+,;，可以考虑做个映射，方便恢复到原始字符串
func SafeFileName(fileName string) string {
	name := path.Clean(path.Base(fileName))
	return fn.ReplaceChars(name, map[rune]rune{
		'\\': '{',
		'/':  '}',
		':':  ';',
		'*':  '+',
		'?':  '@',
		'"':  '`',
		'<':  '[',
		'>':  ']',
		'|':  ',',
		'\n': 0,
		'\r': 0,
	})
}

// SafeFileNameFromURL
// 未测试
func SafeFileNameFromURL(str string) string {
	// PathEscape 对应 decodeURI
	// QueryEscape 对应 decodeURIComponent
	return url.PathEscape(str)
}
