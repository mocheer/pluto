package ds_arcgis_json

import (
	"encoding/json"
	"os"
)

// ReadFile
func ReadFile(fileName string) (*ArcGISJSON, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	v := &ArcGISJSON{}
	err = json.Unmarshal(data, v)
	return v, err
}
