package fs

import (
	"encoding/json"
	"os"
)

// ReadJSON
func ReadJSON(path string, e interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &e)
}

// ReadJSONToMap
func ReadJSONToMap(path string) (data map[string]interface{}, err error) {
	err = ReadJSON(path, data)
	return
}
