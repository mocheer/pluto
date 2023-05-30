package ds_toml

import (
	"bytes"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/mocheer/pluto/pkg/ds"
)

// ReadFile 读取 toml 文件
func ReadFile(fileName string, obj any) error {
	_, err := toml.DecodeFile(fileName, obj)
	return err
}

// Read
func Read(data string, obj any) error {
	_, err := toml.Decode(data, obj)
	return err
}

// Write
func Write(bs *bytes.Buffer, obj any) error {
	err := toml.NewEncoder(bs).Encode(obj)
	if err != nil {
		return err
	}
	return nil
}

// Save
func Save(fileName string, obj any) error {
	var bs bytes.Buffer
	err := Write(&bs, obj)
	if err != nil {
		return err
	}
	fmt.Println(bs.String())
	return ds.Save(fileName, bs.Bytes())
}
