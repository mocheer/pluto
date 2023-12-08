package ds_nc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/batchatco/go-native-netcdf/netcdf"
	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/ds/ds_nc"
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
	t.Log(nc.Attributes())
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
	t.Log(lats)

}

func Test2(t *testing.T) {
	// Open the file
	result, err := ds_nc.ReadFile("./testdata/20230812_155500.nc")
	if err != nil {
		t.Error(err)
	}

	Longitude, _ := result.Data["Latitude"]
	t.Log(Longitude)
	ds.Save("./testdata/Latitude.json", []byte(strings.ReplaceAll(fmt.Sprintf("%v", Longitude), " ", ",")))

}
func Test6(t *testing.T) {
	// 楚志刚@南京大学
	nc, err := netcdf.Open("./testdata/20230812_155500.nc")
	if err != nil {
		panic(err)
	}
	defer nc.Close()
	t.Log(nc.ListVariables()) //列出所有变量名
	t.Log(nc.ListTypes())
	t.Log(nc.ListSubgroups()) //
	t.Log(nc.ListDimensions())
	t.Log(nc.Attributes())
	t.Log(nc.Attributes().Keys())
	t.Log(nc.GetDimension("t"))

	// Read the NetCDF variable from the file

}
