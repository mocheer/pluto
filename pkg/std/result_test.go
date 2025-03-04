package std_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/std"
)

func TestResult(t *testing.T) {
	i := std.NewResult(1, nil)
	t.Log(i)
	t.Log(i.UnWrap())
}
