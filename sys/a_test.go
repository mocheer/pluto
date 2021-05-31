package sys_test

import (
	"testing"

	"github.com/mocheer/pluto/sys"
)

func TestGetObjects(t *testing.T) {
	objs := sys.GetObjects("time")
	t.Error(objs)
}
