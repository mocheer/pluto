package tileset

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

// FromFile
func FromFile(fileName string) (*Tileset, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return FromReader(f), err
}

// FromBytes 从bytes数据中实例化Picture对象
func FromBytes(bs []byte) *Tileset {
	return FromReader(bytes.NewBuffer(bs))
}

func FromReader(data io.Reader) *Tileset {
	var ts *Tileset
	json.NewDecoder(data).Decode(&ts)
	return ts
}
