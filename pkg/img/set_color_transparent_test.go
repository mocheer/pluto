package img_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/img"
)

// TestSetColorTransparent
func TestSetColorTransparent(t *testing.T) {
	p, _, err := img.FromFile("testdata/tile_data.jpg")
	if err != nil {
		t.Error(err)
	}
	i := img.SetColorTransparentWithBlack(p.Image)

	i2 := img.FromImage(i)
	i2.Save("testdata/tile_data_transparent.png", "png")
}
