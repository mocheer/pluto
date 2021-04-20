package fn_test

import (
	"testing"

	"github.com/mocheer/pluto/fn"
)

func TestFmtString(t *testing.T) {
	result := fn.FmtString(`{a}bcd{e}fg{h}`, map[string]interface{}{"a": "1", "b": 2, "c": 3.0, "h": "4.0"})
	if result != "1bcdfg4.0" {
		t.Error(result)
	}
}
