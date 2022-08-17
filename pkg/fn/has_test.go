package fn_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/fn"
	"github.com/stretchr/testify/assert"
)

func TestHas(t *testing.T) {
	type User struct {
		Name string
		Age  uint64
	}

	u1 := User{
		Name: "mocheer",
	}
	u2 := &User{
		Name: "strong",
	}
	assert.Equal(t, fn.Has(u1, "name"), false)
	assert.Equal(t, fn.Has(u2, "name"), false)
	assert.Equal(t, fn.Has(u1, "Name"), true)
	assert.Equal(t, fn.Has(u2, "Name"), true)
}

func TestHas2(t *testing.T) {
	type User struct {
		Name string
		Age  uint64
	}

	u1 := User{
		Name: "mocheer",
	}
	data := fn.GetFieldValueWithLower(u1, "name")
	assert.Equal(t, data.String(), "mocheer")

}
