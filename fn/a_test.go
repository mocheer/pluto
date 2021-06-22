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

func TestUnicode2String(t *testing.T) {
	assert.Equal(t, fn.Unicode2ZH("\u767e\u5ea6\u4e00\u4e0b\uff0c\u4f60\u5c31\u77e5\u9053"), "百度一下，你就知道")
}

// TestRandString 测试生成随机字符串
func TestRandString(t *testing.T) {
	result := fn.RandString(6)
	t.Log(result)
}

func TestToSnakeCase(t *testing.T) {
	result := fn.ToSnakeCase("snakeCase")
	t.Log(result)

	result = fn.ToSnakeCase("CamelcaseToSnakecase")
	t.Log(result)
}
