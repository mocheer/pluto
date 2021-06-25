package calc_test

import (
	"testing"

	"github.com/mocheer/pluto/calc"
)

func TestToWGS84(t *testing.T) {
	lonlats, err := calc.ProjInverse(2633458.580, 470452.902)
	if err != nil {
		t.Log(err)
	}
	t.Log(lonlats)
}
