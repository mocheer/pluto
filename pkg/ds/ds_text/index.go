package ds_text

import (
	"bufio"
	"os"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/fn"
)

// ReadFile
func ReadFile(fileName string) (string, error) {
	bs, err := os.ReadFile(fileName)
	return fn.BytesToString(bs), err
}

// MustReadFile 读取文本文件，当发生错误的时候直接panic
func MustReadFile(fileName string) string {
	return fn.BytesToString(ds.MustReadFile(fileName))
}

// ReadFileByLine 按行读取文本文件，常用于特殊数据文件的解析
func ReadFileByLine(fileName string, fn func(line string)) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fn(scanner.Text())
	}
	return nil
}
