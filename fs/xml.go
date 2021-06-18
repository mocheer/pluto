package fs

import "encoding/xml"

// ReadXML 读取xml文件
func ReadXML(fileName string, e interface{}) error {
	return xml.Unmarshal(MustRead(fileName), &e)
}

// ReadXMLToMap
func ReadXMLToMap(fileName string) (data map[string]interface{}, err error) {
	err = ReadXML(fileName, data)
	return
}
