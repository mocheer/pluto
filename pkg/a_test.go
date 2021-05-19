package pkg_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg"
)

func TestGetObjects(t *testing.T) {
	objs := pkg.GetObjects("time")
	t.Error(objs)
}
