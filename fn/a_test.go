package fn_test

import (
	"reflect"
	"testing"

	"github.com/mocheer/pluto/assert"
	"github.com/mocheer/pluto/fn"
)

// TestFmtString fn.FmtString
func TestFmtString(t *testing.T) {
	result := fn.FmtString(`{a}bcd{e}fg{h}`, map[string]interface{}{"a": "1", "b": 2, "c": 3.0, "h": "4.0"})
	assert.Equal(t, result, "1bcdfg4.0")
}

func TestGetType(t *testing.T) {
	assert.Equal(t, fn.GetType("string"), reflect.String)
	assert.Equal(t, fn.GetType(1), reflect.Int)
	assert.Equal(t, fn.GetType(nil), reflect.Invalid)
}
