package gen_test

import (
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/mocheer/pluto/pkg/gen"
)

func TestGen(t *testing.T) {
	m := gen.Package("name")
	t.Log(m.String())
}

func TestGen2(t *testing.T) {
	m := gen.Package("name")
	m.AddVar("a", jen.Lit("string value"))
	m.AddVar("b", jen.Lit(5))
	t.Log(m.String())
}
