package img_test

import (
	"testing"

	"github.com/mocheer/pluto/ts/img"
)

// TestSave
func TestSave(t *testing.T) {
	p, err := img.FromFile("test.png")
	if err != nil {
		t.Error(err)
	}
	p.Save("test_save.png")
}
