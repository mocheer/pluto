package ds_xml

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"

	"github.com/mocheer/pluto/pkg/ds"
)

// ReadFile 读取xml文件
func ReadFile(fileName string, e any) error {
	return xml.Unmarshal(ds.MustReadFile(fileName), &e)
}

// ReadFileToDocument
func ReadFileToDocument(fileName string) (*Document, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ReadToDocument(file)
}

// ReadBufferToDocument
func ReadBufferToDocument(buffer []byte) (*Document, error) {
	return ReadToDocument(bytes.NewReader(buffer))
}

// ReadToDocument 读取xml文件，返回解析后的文档
func ReadToDocument(reader io.Reader) (*Document, error) {
	return ParseToTree(reader)
}
