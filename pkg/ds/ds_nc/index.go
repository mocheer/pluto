package ds_nc

import (
	"github.com/batchatco/go-native-netcdf/netcdf"
	"github.com/batchatco/go-native-netcdf/netcdf/api"
)

type DsNc struct {
	Data  map[string]any
	Attrs map[string]any
}

func Read(file api.ReadSeekerCloser) *DsNc {
	nc, err := netcdf.New(file)
	if err != nil {
		panic(err)
	}
	defer nc.Close()
	return readNc(nc)
}

func ReadFile(fileName string) (*DsNc, error) {
	nc, err := netcdf.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer nc.Close()
	return readNc(nc), nil
}

// readNc
// TODO 读取变量
func readNc(nc api.Group) *DsNc {
	data := map[string]any{}
	keys := nc.ListVariables()
	for _, key := range keys {
		val, _ := nc.GetVariable(key)
		data[key] = val.Values
	}
	return &DsNc{Data: data}
}
