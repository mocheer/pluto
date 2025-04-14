package gast

import (
	"testing"

	"github.com/mocheer/xena/pkg/gm"
)

func TestF(t *testing.T) {
	code, err := ExtractFuncCode("./a_test.go", "TestF")
	if err != nil {
		t.Log(err)
	}
	t.Log(code)
	// t.Log(strings.TrimSpace(code))
}

func TestMethod(t *testing.T) {
	m, err := GetModelMethod((gm.Point{}).ToLonLat)
	if err != nil {
		panic(err)
	}
	t.Log(m.currentFile)
	t.Logf("%+v", m.Methods[0])
}

func TestGetImports(t *testing.T) {
	t.Log(GetImports("get_imports.go"))
	t.Log(GetFunctionImports("get_imports.go", "GetFunctionImports"))
}
