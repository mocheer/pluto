package ds_shp

import (
	"fmt"
	"log"
	"reflect"

	"github.com/jonas-p/go-shp"
)

func ReadFile(fileName string) {
	// open a shapefile for reading
	shape, err := shp.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer shape.Close()

	// fields from the attribute table (DBF)
	fields := shape.Fields()

	// loop through all features in the shapefile
	for shape.Next() {
		n, p := shape.Shape()

		// print feature
		fmt.Println(reflect.TypeOf(p).Elem(), p.BBox())

		// print attributes
		for k, f := range fields {
			val := shape.ReadAttribute(n, k)
			fmt.Printf("\t%v: %v\n", f, val)
		}

	}
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
