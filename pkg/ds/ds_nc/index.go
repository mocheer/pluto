package ds_nc

import (
	"github.com/batchatco/go-native-netcdf/netcdf"
	"github.com/batchatco/go-native-netcdf/netcdf/api"
)

type DsNcOptions struct {
	LonField   string
	LatField   string
	ValueField string
}

type DsNc struct {
	Lons   []float32
	Lats   []float32
	Values []float32
	Attrs  map[string]any
}

func Read(file api.ReadSeekerCloser, options DsNcOptions) *DsNc {
	nc, err := netcdf.New(file)
	if err != nil {
		panic(err)
	}
	defer nc.Close()
	return readNc(nc, options)
}

func ReadFile(fileName string, options DsNcOptions) *DsNc {

	nc, err := netcdf.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer nc.Close()
	return readNc(nc, options)
}

func readNc(nc api.Group, options DsNcOptions) *DsNc {
	vr, _ := nc.GetVariable(options.LonField)
	if vr == nil {
		panic("lon variable not found")
	}
	lons, _ := vr.Values.([]float32)
	//
	vr2, _ := nc.GetVariable(options.LatField)
	if vr2 == nil {
		panic("lat variable not found")
	}
	lats, _ := vr2.Values.([]float32)
	//
	vr3, _ := nc.GetVariable(options.ValueField)
	if vr3 == nil {
		panic("value variable not found")
	}
	values, _ := vr3.Values.([]float32)
	return &DsNc{
		Lons:   lons,
		Lats:   lats,
		Values: values,
	}
}
