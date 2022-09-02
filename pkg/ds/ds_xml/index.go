package ds_xml

import (
	"encoding/xml"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/ts"
)

// Read 读取xml文件
func Read(fileName string, e any) error {
	return xml.Unmarshal(ds.MustReadFile(fileName), &e)
}

// ReadToMap
func ReadToMap(fileName string) (data ts.Map[any], err error) {
	err = Read(fileName, data)
	return
}
