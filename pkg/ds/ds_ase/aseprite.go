package ds_ase

import (
	"encoding/json"
	"os"
)

// ReadJSON 读取aseprite的json文件
func ReadJSON(f string) (*Aseprite, error) {
	// read json
	data, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}
	aseprite := &Aseprite{}
	err = json.Unmarshal(data, aseprite)
	if err != nil {
		return nil, err
	}
	return aseprite, nil
}
