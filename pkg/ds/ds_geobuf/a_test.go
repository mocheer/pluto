package ds_geobuf_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/ds/ds_geobuf"
	"github.com/spatial-go/geoos/geoencoding/geobuf"
)

func Test(t *testing.T) {
	bs := ds.MustReadFile("./testdata/featurecollection.geobuf")
	encode := &geobuf.GeobufEncoder{}
	encode.Decode(bs)

}

func Test2(t *testing.T) {
	j, _ := ds_geobuf.Read("./testdata/featurecollection.geobuf").ToGeoJSON().MarshalJSON()
	t.Log(string(j))
}

func Test3(t *testing.T) {
	j, _ := ds_geobuf.Read("./testdata/feature.geobuf").ToGeoJSON().MarshalJSON()
	t.Log(string(j))
}
