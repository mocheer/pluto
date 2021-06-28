package sys_test

import (
	"testing"

	"github.com/mocheer/pluto/sys"
)

func TestGetObjects(t *testing.T) {
	objs := sys.GetObjects("testing")
	t.Log(objs)
}

func TestPowerShell(t *testing.T) {
	var s sys.PowerShell = "npm -v\n"
	s.Run()
}
