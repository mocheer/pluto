package ds_shp

import (
	"archive/zip"
	"io"
	"log"

	"github.com/mocheer/pluto/pkg/ds/ds_shp/go-shp"
	"github.com/mocheer/xena/pkg/gm"
	"github.com/mocheer/xena/pkg/gm/geojson"
	"github.com/samber/lo"
)

// ReadFile
// ReadFile("xx.shp")
func ReadFile(fileName string) *geojson.GeometryCollection {
	// open a shapefile for reading
	shape, err := shp.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer shape.Close()
	return toGeoJSON(shape)
}

type shpReader interface {
	Fields() []shp.Field
	Next() bool
	Shape() (int, shp.Shape)
	// ReadAttribute(row int, field int) string
	Attribute(field int) string
}

func toGeoJSON(shape shpReader) *geojson.GeometryCollection {
	data := geojson.NewGeometryCollection()
	// fields from the attribute table (DBF)
	// 这里是读取同一路径下的dbf文件
	fields := shape.Fields()

	//
	for shape.Next() {
		_, p := shape.Shape() // 忽略 rowNum
		// log.Println(shape.GeometryType)
		var geom *geojson.Geometry
		switch g := p.(type) {
		case *shp.Point:
			poly := gm.Point{g.X, g.Y}.SetPrecision(6)
			geom = geojson.NewPointGeom(&poly)
		case *shp.Polygon:
			poly := make(gm.Polygon, 1)
			poly[0] = lo.Map(g.Points, func(p shp.Point, _ int) [2]float64 {
				return gm.Point{p.X, p.Y}.SetPrecision(6)
			})
			geom = geojson.NewPolygonGeom(&poly)
		case *shp.PolyLine:
			var poly gm.LineString = lo.Map(g.Points, func(p shp.Point, _ int) [2]float64 {
				return gm.Point{p.X, p.Y}.SetPrecision(6)
			})
			geom = geojson.NewLineStringGeom(&poly)
		}
		data.Append(geom)
		// 读取数据， k为key，f为值
		geom.Properties = geojson.Properties{}
		for k, f := range fields {
			// val := shape.ReadAttribute(n, k)
			val := shape.Attribute(k)
			geom.Properties[f.String()] = val
		}
	}
	return data
}

// go-shp 暂不支持
func Read(r io.ReaderAt, size int64) *geojson.GeometryCollection {
	zr, err := zip.NewReader(r, size)
	if err != nil {
		return nil
	}
	shapeZip, err := shp.OpenZipReader(zr, nil)
	if err != nil {
		return nil
	}
	defer shapeZip.Close()
	return toGeoJSON(shapeZip)
}

func Write() {
	// // points to write
	// points := []shp.Point{
	// 	shp.Point{10.0, 10.0},
	// 	shp.Point{10.0, 15.0},
	// 	shp.Point{15.0, 15.0},
	// 	shp.Point{15.0, 10.0},
	// }

	// // fields to write
	// fields := []shp.Field{
	// 	// String attribute field with length 25
	// 	shp.StringField("NAME", 25),
	// }

	// // create and open a shapefile for writing points
	// shape, err := shp.Create("points.shp", shp.POINT)
	// if err != nil { log.Fatal(err) }
	// defer shape.Close()

	// // setup fields for attributes
	// shape.SetFields(fields)

	// // write points and attributes
	// for n, point := range points {
	// 	shape.Write(&point)

	// 	// write attribute for object n for field 0 (NAME)
	// 	shape.WriteAttribute(n, 0, "Point " + strconv.Itoa(n + 1))
	// }
}
