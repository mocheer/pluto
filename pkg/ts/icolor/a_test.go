package icolor_test

import (
	"image/color"
	"testing"

	"github.com/mocheer/pluto/pkg/ts/icolor"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	c := color.RGBA{255, 0, 0, 255}
	assert.Equal(t, icolor.ToHex(c), "FF0000FF")
}

func Test2(t *testing.T) {
	rgba := "rgba(255, 0, 0, 0.5)"
	c := icolor.FromRGBAString(rgba)
	assert.Equal(t, icolor.ToHex(c), "FF00007F")
}

func Test3(t *testing.T) {
	rgba := "rgba(255, 0, 0, 0.5)"
	c := icolor.FromHtmlColor(rgba)
	assert.Equal(t, icolor.ToHex(c), "FF00007F")
}
