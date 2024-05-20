package ds_json

import (
	"encoding/json"
	"os"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/fn"

	"github.com/tidwall/gjson"
)

// ReadFile
func ReadFile(fileName string, e any) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, e)
}

// ReadGJSON
func ReadGJSON(fileName string) gjson.Result {
	return gjson.Parse(fn.BytesToString(ds.MustReadFile(fileName)))
}

// Save 保存数据为json文件
func Save(fileName string, value any) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	f, err := ds.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	return err
}

// SaveWithIndent 保存数据为json文件（包含json格式化）
func SaveWithIndent(fileName string, value any) error {
	data, err := json.MarshalIndent(value, "", " ")
	if err != nil {
		return err
	}
	f, err := ds.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	return err
}
