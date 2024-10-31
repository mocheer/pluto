package ds

import (
	"io"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/mocheer/pluto/pkg/fn"
)

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

func Basename(name string) string {
	// 获取基础文件名，包括后缀
	baseName := filepath.Base(name)
	// 获取最后一个点（.）的位置
	if lastDot := strings.LastIndexByte(baseName, '.'); lastDot != -1 {
		// 截取不带后缀的文件名
		baseName = baseName[:lastDot]
	}
	return baseName
}
