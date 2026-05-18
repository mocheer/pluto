package atlas_test

import (
	"testing"

	atlas "github.com/mocheer/pluto/pkg/ds/ds_atlas"
)

func TestRead(t *testing.T) {
	a, err := atlas.Read("testdata/common.altas")
	if err != nil {
		t.Fatal(err)
	}
	err = a.Save("testdata/common")
	if err != nil {
		t.Fatal(err)
	}
}
