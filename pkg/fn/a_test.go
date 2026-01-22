package fn_test

import (
	"fmt"
	"math"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/fn"
	"github.com/stretchr/testify/assert"
)

func TestMD(t *testing.T) {
	data, err := os.ReadFile("./testdata/LAND STALKER.md")
	if err != nil {
		panic(err)
	}

	for i, v := range data {
		switch v {
		case 0xff:
			if data[i+1] == 0xd8 {

				for j, v := range data[i:] {
					if v == 0xff && data[i+j+1] == 0xD9 {
						t.Log(i, i+j, j)
						ds.Save(fmt.Sprintf("testdata/part/%d.jpg", i), data[i:i+j+2])
						break
					}
				}

			}
		}
	}

}

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
func TestBase64(t *testing.T) {
	data, _ := ds.ReadFile("testdata/auth.jpg")
	// t.Log(fn.BtoaBytes([]byte{0xFF, 0xD8}))
	t.Log(data[0:2])
	t.Log(fn.BtoaBytes(data[0:2]))
	t.Log(fn.BtoaBytes(data[0:3]))
}

func TestGo(t *testing.T) {
	fn.Go(
		func() {
			t.Log("1")
		},
		func() {
			time.Sleep(time.Second)
			t.Log("2")
		},
		func() {
			t.Log("3")
		},
	)
	t.Log("4")
}
