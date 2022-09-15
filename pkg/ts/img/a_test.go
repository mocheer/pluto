package img_test

import (
	"image/color"
	"strconv"
	"testing"

	"github.com/mocheer/pluto/pkg/ts/img"
)

// TestSave
func TestSave(t *testing.T) {
	p, err := img.FromFile("test.png")
	if err != nil {
		t.Error(err)
	}
	p.Save("test_save.png")
}

// TestColor
func TestColor(t *testing.T) {
	c := color.RGBA{255, 16, 16, 254}
	r, g, b, a := c.RGBA()
	t.Log(r, g, b, a)
	var hexr, hexg, hexb string
	hexr = strconv.FormatUint(uint64(r), 16)
	if r < 16 {
		hexr = "0" + hexr
	}
	hexg = strconv.FormatUint(uint64(g), 16)
	if g < 16 {
		hexg = "0" + hexg
	}
	hexb = strconv.FormatUint(uint64(b), 16)
	if b < 16 {
		hexb = "0" + hexb
	}

	t.Log(hexr + hexg + hexb)
}
