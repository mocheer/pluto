package ds_ini

import "gopkg.in/ini.v1"

// Read
func Read(fileName string, e interface{}) error {
	return ini.MapTo(e, fileName)
}
