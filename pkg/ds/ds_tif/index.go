package ds_tif

import (
	"bytes"
	"os"

	"github.com/google/tiff"
)

func ReadFile(fileName string) (*DsTif, error) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return parse(f)
}

func ReadBytes(bs []byte) (*DsTif, error) {
	f := bytes.NewReader(bs)
	return parse(f)
}

func parse(reader tiff.ReadAtReadSeeker) (*DsTif, error) {
	tif, err := tiff.Parse(reader, nil, nil)
	if err != nil {
		panic(err)
	}
	m := &DsTif{Tif: tif}
	// m.IFD_Index = len(m.Tif.IFDs()) - 1
	m.Data, err = m.readData()
	return m, err
}
