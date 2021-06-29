package sys_test

import (
	"testing"

	"github.com/mocheer/pluto/sys"
)

func TestGetObjects(t *testing.T) {
	objs := sys.GetObjects("testing")
	t.Log(objs)
}

func TestExec(t *testing.T) {
	sys.Exec("npm", "-v")
}

func TestShell(t *testing.T) {
	sys.Shell("node -v;npm -v")
}
