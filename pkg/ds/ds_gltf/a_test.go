package ds_gltf_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/ds/ds_gltf"
)

func TestRead(t *testing.T) {
	doc, err := ds_gltf.ReadFile("./testdata/phoenix.glb")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", doc.Asset)
}
