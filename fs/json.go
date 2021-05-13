package fs

import (
	"encoding/json"
)

// ReadJSON
func ReadJSON(path string, e interface{}) error {
	return json.Unmarshal(MustReadFile(path), &e)
}

// ReadJSONToMap
func ReadJSONToMap(path string) (data map[string]interface{}, err error) {
	err = ReadJSON(path, data)
	return
}

// SaveJSON 保存为json文件
func SaveJSON(path string, e interface{}) error {
	data, err := json.MarshalIndent(e, "", " ")
	if err != nil {
		return err
	}
	file, err := Create(path)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	return err
}
