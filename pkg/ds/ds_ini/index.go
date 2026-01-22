package ds_ini

import "gopkg.in/ini.v1"

// ReadFile
func ReadFile(fileName string, e any) error {
	return ini.MapTo(e, fileName)
}
