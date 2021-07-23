package ds

import (
	"io"
	"os"
	"strings"
)

// Append 往文件尾部添加字符串
func Append(fileName string, content string) {
	f, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	f.Write([]byte(content))
	f.Close()
}

// Append 往文件尾部添加字符串
func AppendHead(fileName string, content string) {
	f, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	defer f.Close()
	data, err := io.ReadAll(f)
	old := string(data)
	if err == nil && !strings.Contains(old, content) {
		f.WriteAt([]byte(content+"\n"+old), 0)
	}
}
