package dsxml

import (
	"encoding/xml"

	"github.com/mocheer/pluto/ds"
	"github.com/mocheer/pluto/ts"
)

// Read 读取xml文件
func Read(fileName string, e interface{}) error {
	return xml.Unmarshal(ds.MustRead(fileName), &e)
}

// ReadToMap
func ReadToMap(fileName string) (data ts.Map, err error) {
	err = Read(fileName, data)
	return
}
