package fn_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/fn"
)

func TestNanoid(t *testing.T) {
	t.Log(fn.Nanoid(8))
}
