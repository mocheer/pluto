package fs

import (
	"encoding/json"

	"github.com/mocheer/pluto/fn"
	"github.com/mocheer/pluto/js/JSON"
	"github.com/mocheer/pluto/ts"
	"github.com/tidwall/gjson"
)

// ReadJSON
func ReadJSON(fileName string, e interface{}) error {
	return json.Unmarshal(MustRead(fileName), e)
}

// ReadJSONToMap
func ReadJSONToMap(fileName string) (data ts.Map, err error) {
	err = ReadJSON(fileName, &data)
	return
}

// ReadJSONToGJSON
func ReadJSONToGJSON(fileName string) gjson.Result {
	return JSON.Parse(fn.B2S(MustRead(fileName)))
}

// SaveJSON 保存为json文件
func SaveJSON(fileName string, e interface{}) error {
	data, err := json.MarshalIndent(e, "", " ")
	if err != nil {
		return err
	}
	f, err := Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	return err
}
