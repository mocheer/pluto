package fn_test

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/mocheer/pluto/pkg/fn"
	"github.com/stretchr/testify/assert"
)

func TestB2S(t *testing.T) {
	var b = []byte(`b2s`)
	var s = fn.BytesToString(b)
	var s2 = strings.Clone(s)
	b[1] = 1

	fmt.Println(s, s2)
}

func TestFmtString(t *testing.T) {
	result := fn.Format(`{a}bcd{e}fg{h}`, map[string]any{"a": "1", "b": 2, "c": 3.0, "h": "4.0"})
	assert.Equal(t, result, "1bcdfg4.0")

}

func TestToCamelCase(t *testing.T) {
	assert.Equal(t, fn.ToCamelCase("camel-case"), "camelCase")
}

func TestToSnakeCase(t *testing.T) {
	result := fn.ToSnakeCase("snakeCase")
	t.Log(result)

	result = fn.ToSnakeCase("CamelcaseToSnakecase")
	t.Log(result)
}

func TestMinFloat64(t *testing.T) {
	var a = 4.8
	t.Log(math.SmallestNonzeroFloat64) //struct
	t.Log(int(a))
	t.Log(fn.RoundInt(4.49))
	t.Log(fn.RoundInt(4.5))
}
