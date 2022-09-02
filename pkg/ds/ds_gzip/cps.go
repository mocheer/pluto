package ds_gzip

import (
	"bytes"
	"compress/gzip"
	"io"
)

// Encode 读取 byte 数据 写入并返回 gzip 字节数组
func Encode(data []byte) []byte {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()
	w.Write(data)
	w.Flush()
	return b.Bytes()
}

// Decode 读取 gzip的字节数组，返回原始的byte数据
func Decode(data []byte) ([]byte, error) {
	dr := bytes.NewReader(data)
	gr, err := gzip.NewReader(dr)
	if err != nil {
		return nil, err
	}
	defer gr.Close()
	buf, err := io.ReadAll(gr)
	if err != nil { // 会出现unexpected EOF的错误，但buf仍能正常
		return buf, err
	}
	return buf, nil
}
