package fn_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/fn"
)

func TestInt2IP(t *testing.T) {
	ip := fn.IntToIP(3232235521)
	t.Log(ip.String())
}
