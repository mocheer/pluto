package ds_json

import (
	"encoding/json"
	"os"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/fn"

	"github.com/tidwall/gjson"
)

// Read
func Read(fileName string, e any) error {
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

// Save 保存数据为json文件（包含json格式化）
func Save(data any, fileName string) error {
	_data, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	f, err := ds.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(_data)
	return err
}
