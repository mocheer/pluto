package dsjson

import (
	"encoding/json"

	"github.com/mocheer/pluto/ds"
	"github.com/mocheer/pluto/fn"

	"github.com/mocheer/pluto/jsg/JSON"
	"github.com/tidwall/gjson"
)

// Read
func Read(fileName string, e interface{}) error {
	data, err := ds.Read(fileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, e)
}

// ReadGJSON
func ReadGJSON(fileName string) gjson.Result {
	return JSON.Parse(fn.B2S(ds.MustRead(fileName)))
}

// Save 保存为json文件
func Save(fileName string, e interface{}) error {
	data, err := json.MarshalIndent(e, "", " ")
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
