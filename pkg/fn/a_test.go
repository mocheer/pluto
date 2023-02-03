package fn_test

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"testing"
	"time"

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
	result := fn.FmtString(`{a}bcd{e}fg{h}`, map[string]any{"a": "1", "b": 2, "c": 3.0, "h": "4.0"})
	assert.Equal(t, result, "1bcdfg4.0")

}

func TestGetKind(t *testing.T) {
	assert.Equal(t, fn.GetKind(1), reflect.Int)
	assert.Equal(t, fn.GetKind(1.0), reflect.Float64)
	assert.Equal(t, fn.GetKind(nil), reflect.Invalid)
	assert.Equal(t, fn.GetKind("string"), reflect.String)
	assert.Equal(t, fn.GetKind(map[string]string{}), reflect.Map)
	assert.Equal(t, fn.GetKind(map[string]int{}), reflect.Map)
	assert.Equal(t, fn.GetKind(false), reflect.Bool)
}

func TestToCamelCase(t *testing.T) {
	assert.Equal(t, fn.ToCamelCase("camel-case"), "camelCase")
}

func TestUnicode2String(t *testing.T) {
	assert.Equal(t, fn.Unicode2ZH("\u767e\u5ea6\u4e00\u4e0b\uff0c\u4f60\u5c31\u77e5\u9053"), "百度一下，你就知道")
}

func TestToSnakeCase(t *testing.T) {
	result := fn.ToSnakeCase("snakeCase")
	t.Log(result)

	result = fn.ToSnakeCase("CamelcaseToSnakecase")
	t.Log(result)
}

func TestIsTime(t *testing.T) {

	t.Log(fn.GetType(time.Now())) //struct
}

func TestMinFloat64(t *testing.T) {
	var a = 4.8
	t.Log(math.SmallestNonzeroFloat64) //struct
	t.Log(int(a))
	t.Log(fn.RoundInt(4.49))
	t.Log(fn.RoundInt(4.5))
}
