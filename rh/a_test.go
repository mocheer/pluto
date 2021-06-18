package rh_test

import (
	"testing"

	"github.com/mocheer/pluto/rh"
)

func TestGet(t *testing.T) {
	data, _ := rh.Get("https://www.baidu.com/")
	t.Log(string(data))
}
