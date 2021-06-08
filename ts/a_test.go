package ts_test

import (
	"testing"

	"github.com/mocheer/pluto/assert"
	"github.com/mocheer/pluto/ts"
)

func TestGetObjects(t *testing.T) {

	c := ts.ColorFromCSS("#FF0000")
	assert.Equal(t, c.R, uint8(255))
	assert.Equal(t, c.G, uint8(0))
	assert.Equal(t, c.B, uint8(0))
	assert.Equal(t, c.A, uint8(255))

	c = ts.ColorFromCSS("#fff")
	assert.Equal(t, c.R, uint8(255))
	assert.Equal(t, c.G, uint8(255))
	assert.Equal(t, c.B, uint8(255))
	assert.Equal(t, c.A, uint8(255))

	c = ts.ColorFromCSS("red")
	assert.Equal(t, c.R, uint8(255))
	assert.Equal(t, c.G, uint8(0))
	assert.Equal(t, c.B, uint8(0))
	assert.Equal(t, c.A, uint8(255))

	c = ts.ColorFromCSS("rgba(255,0,0,1)")
	assert.Equal(t, c.R, uint8(255))
	assert.Equal(t, c.G, uint8(0))
	assert.Equal(t, c.B, uint8(0))
	assert.Equal(t, c.A, uint8(255))

	c = ts.ColorFromCSS("rgb(255,0,0)")
	assert.Equal(t, c.R, uint8(255))
	assert.Equal(t, c.G, uint8(0))
	assert.Equal(t, c.B, uint8(0))
	assert.Equal(t, c.A, uint8(255))
}
