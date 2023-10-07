package object_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/ts/object"
)

func TestIsString(t *testing.T) {
	type A string
	var b A
	b = "c"
	var c any
	c = b
	d := "d"
	var e any
	e = d
	t.Log(object.IsString(b))
	t.Log(object.IsString(c))
	t.Log(c.(A))
	_, ok := e.(A)
	t.Log(ok)
	t.Log(A(d))
}
