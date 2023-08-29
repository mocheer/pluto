package ds_webp_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/ds/ds_webp"
	"github.com/mocheer/pluto/pkg/ts/img"
)

func TestRead(t *testing.T) {

	buffer, _ := ds_webp.FromImageFile("./testdata/auth.png")
	i, _ := img.FromFile("./testdata/auth.png")
	img.SaveAsJPEG(i.Image, "./testdata/auth.jpg")
	ds.Save("./testdata/auth.webp", buffer.Bytes())

}
