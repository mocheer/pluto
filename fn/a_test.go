package fn_test

import (
	"reflect"
	"testing"

	"github.com/mocheer/pluto/assert"
	"github.com/mocheer/pluto/fn"
)

func TestFmtString(t *testing.T) {
	result := fn.FmtString(`{a}bcd{e}fg{h}`, map[string]interface{}{"a": "1", "b": 2, "c": 3.0, "h": "4.0"})
	assert.Equal(t, result, "1bcdfg4.0")
}

func TestGetKind(t *testing.T) {
	assert.Equal(t, fn.GetKind("string"), reflect.String)
	assert.Equal(t, fn.GetKind(1), reflect.Int)
	assert.Equal(t, fn.GetKind(nil), reflect.Invalid)
}

func TestToCamelCase(t *testing.T) {
	assert.Equal(t, fn.ToCamelCase("camel-case"), "camelCase")
}
