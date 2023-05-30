package ds_nc

import (
	"testing"

	"github.com/batchatco/go-native-netcdf/netcdf"
)

// http://58.59.29.50:11011/jxyy/analytical/
func Test1(t *testing.T) {
	// Open the file
	nc, err := netcdf.Open("./testdata/58.59.29.50.11011.nc")
	if err != nil {
		panic(err)
	}
	defer nc.Close()
	t.Log(nc.ListVariables()) //列出所有变量名
	t.Log(nc.ListTypes())
	t.Log(nc.ListSubgroups()) //
	t.Log(nc.ListDimensions())
	// Read the NetCDF variable from the file

	vr, _ := nc.GetVariable("lat")
	if vr == nil {
		panic("latitude variable not found")
	}

	// Cast the data into a Go type we can use
	lats, has := vr.Values.([]float32)
	if !has {
		panic("latitude data not found")
	}
	for i, lat := range lats {
		t.Log(i, lat)
	}
}
