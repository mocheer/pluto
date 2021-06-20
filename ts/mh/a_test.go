package mh_test

import (
	"testing"

	"github.com/mocheer/pluto/ts/mh"
)

func TestMH(t *testing.T) {
	data := mh.New("## markdown document").HTML()
	t.Log(data)
}
