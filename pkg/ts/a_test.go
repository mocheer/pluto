package ts_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/assert"
	"github.com/mocheer/pluto/pkg/ts"
)

func TestGetObjects(t *testing.T) {

	c := ts.NewColorFromCSS("#FF0000")
	assert.Equal(t, c.R, uint8(255))
	assert.Equal(t, c.G, uint8(0))
	assert.Equal(t, c.B, uint8(0))
	assert.Equal(t, c.A, uint8(255))

	c = ts.NewColorFromCSS("#fff")
	assert.Equal(t, c.R, uint8(255))
	assert.Equal(t, c.G, uint8(255))
	assert.Equal(t, c.B, uint8(255))
	assert.Equal(t, c.A, uint8(255))

	c = ts.NewColorFromCSS("red")
	assert.Equal(t, c.R, uint8(255))
	assert.Equal(t, c.G, uint8(0))
	assert.Equal(t, c.B, uint8(0))
	assert.Equal(t, c.A, uint8(255))

	c = ts.NewColorFromCSS("rgba(255,0,0,1)")
	assert.Equal(t, c.R, uint8(255))
	assert.Equal(t, c.G, uint8(0))
	assert.Equal(t, c.B, uint8(0))
	assert.Equal(t, c.A, uint8(255))

	c = ts.NewColorFromCSS("rgb(255,0,0)")
	assert.Equal(t, c.R, uint8(255))
	assert.Equal(t, c.G, uint8(0))
	assert.Equal(t, c.B, uint8(0))
	assert.Equal(t, c.A, uint8(255))
}
