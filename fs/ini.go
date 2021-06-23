package fs

import "gopkg.in/ini.v1"

// ReadIni
func ReadIni(fileName string, e interface{}) error {
	return ini.MapTo(e, fileName)
}
