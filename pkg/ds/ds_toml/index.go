package ds_toml

import (
	"bytes"

	"github.com/BurntSushi/toml"
	"github.com/mocheer/pluto/pkg/ds"
)

// ReadFile 读取 toml 文件
func ReadFile(fileName string, e any) error {
	_, err := toml.DecodeFile(fileName, e)
	return err
}

// Read
func Read(data string, e any) error {
	_, err := toml.Decode(data, e)
	return err
}

// Write
func Write(bs bytes.Buffer, e any) error {
	err := toml.NewEncoder(&bs).Encode(e)
	if err != nil {
		return err
	}
	return nil
}

// Save
func Save(fileName string, e any) error {
	var bs bytes.Buffer
	err := Write(bs, e)
	if err != nil {
		return err
	}
	return ds.Save(fileName, bs.Bytes())
}
