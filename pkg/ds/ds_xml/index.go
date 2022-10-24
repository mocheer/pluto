package ds_xml

import (
	"encoding/xml"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/ts"
)

// ReadFile 读取xml文件
func ReadFile(fileName string, e any) error {
	return xml.Unmarshal(ds.MustReadFile(fileName), &e)
}

// ReadFileToMap
func ReadFileToMap(fileName string) (data ts.Map[any], err error) {
	err = ReadFile(fileName, data)
	return
}
