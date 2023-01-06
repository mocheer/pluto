package ds_gltf

import (
	"bytes"
	"os"

	"github.com/qmuntal/gltf"
)

// Read
func Read(data []byte) (gltf.Document, error) {
	var doc gltf.Document
	dr := bytes.NewReader(data)
	gltf.NewDecoder(dr).Decode(&doc)
	return doc, nil
}

// ReadFile
func ReadFile(fileName string) (gltf.Document, error) {
	var doc gltf.Document
	//
	fr, err := os.Open(fileName)
	if err != nil {
		return doc, err
	}
	// defer: 在函数退出时,执行关闭文件
	defer fr.Close()
	gltf.NewDecoder(fr).Decode(&doc)
	return doc, nil
}

// ToBuffer
func ToBuffer(doc gltf.Document, asBubary bool) bytes.Buffer {
	var buf bytes.Buffer
	enc := gltf.NewEncoder(&buf)
	// 默认 false 生成 gltf
	// 设置为true时生成 glb
	enc.AsBinary = asBubary
	enc.Encode(&doc)
	return buf
}
