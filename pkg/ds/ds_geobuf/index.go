package ds_geobuf

import (
	"github.com/murphy214/geobuf"
	geojson "github.com/paulmach/go.geojson"
)

// Geobuf 目前无法正确解析postgis的geobuf数据
type Geobuf struct {
	*geobuf.Reader
}

// ReadFile
func ReadFile(filename string) *Geobuf {
	return &Geobuf{geobuf.ReaderFile(filename)}
}

// func Reader(file *os.File) *Geobuf {
// 	reader := bufio.NewReader(file)
// 	buf := &geobuf.Reader{
// 		Reader:   protoscan.NewProtobufScanner(reader),
// 		Filename: file.Name(),
// 		FileBool: true,
// 		File:     file,
// 	}
// 	buf.CheckMetaData()
// 	buf.FeatureCount = 0
// 	return &Geobuf{buf}
// }

// Read
func Read(data []byte) *Geobuf {
	return &Geobuf{geobuf.ReaderBuf(data)}
}

// ToGeoJSON
func (m Geobuf) ToGeoJSON() *geojson.FeatureCollection {
	fc := geojson.NewFeatureCollection()
	for m.Next() {
		fc.AddFeature(m.Feature())
	}
	return fc
}
