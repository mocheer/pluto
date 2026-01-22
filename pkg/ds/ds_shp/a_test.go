package ds_shp_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/ds/ds_json"
	"github.com/mocheer/pluto/pkg/ds/ds_shp"
)

func TestShp2Geojson(t *testing.T) {
	data := ds_shp.ReadFile("./testdata/金水河流域过程数据/grid_jsh_5m_1.shp")
	ds_json.SaveWithIndent("./testdata/geojson.json", data)
}
