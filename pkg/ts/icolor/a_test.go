package icolor_test

import (
	"image/color"
	"testing"

	"github.com/mocheer/pluto/pkg/ts/icolor"
)

func Test1(t *testing.T) {
	c := color.RGBA{255, 0, 0, 255}
	t.Log(icolor.ToHex(c))
}
