package fs

import (
	"encoding/json"
)

// ReadJSON
func ReadJSON(fileName string, e interface{}) error {
	return json.Unmarshal(MustReadFile(fileName), &e)
}

// ReadJSONToMap
func ReadJSONToMap(fileName string) (data map[string]interface{}, err error) {
	err = ReadJSON(fileName, data)
	return
}

// SaveJSON 保存为json文件
func SaveJSON(fileName string, e interface{}) error {
	data, err := json.MarshalIndent(e, "", " ")
	if err != nil {
		return err
	}
	file, err := Create(fileName)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	return err
}
