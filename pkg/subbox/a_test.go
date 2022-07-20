package subbox_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/subbox"
)

func TestRun(t *testing.T) {
	vm := subbox.New()
	val, err := vm.Run("var module ={};module.exports = {a:1};")
	if err != nil {
		t.Error(err)
	}

	t.Log(val.Export())
	t.Log(vm.Run("JSON.stringify(module)"))

}

func TestPanic(t *testing.T) {
	vm := subbox.New()
	defer func() {
		recover()
		t.Error("触发了panic")
	}()
	vm.Set("panic", func(msg string) {
		t.Log(msg)
		panic(msg)
	})
	t.Log(vm.Run("typeof panic"))
	t.Log(vm.Run("panic('test panic')"))
}

func TestEnv(t *testing.T) {
	vm := subbox.New()
	moduleVal, err := vm.Run("typeof exports === 'object' && typeof module !== 'undefined'")
	if err != nil {
		t.Error(err)
	}
	t.Log(moduleVal.Export())
	//
	thisVal, err := vm.Run("this")
	if err != nil {
		t.Error(err)
	}
	t.Log(thisVal.Export())

	//
	windowVal, err := vm.Run("window")
	if err != nil {

	}

	t.Log(windowVal)

	t.Log(vm.Global().Export())
	vm.Run("this.a = 1")
	vm.Run("this.b = function(){ return a+1 }")
	t.Log(vm.Global().Export())

	t.Log(vm.Get("b"))

	b2, err := vm.Call("b", nil)

	t.Log(err)
	t.Log(b2)

	//
	t.Error("")
}

func TestGetSet(t *testing.T) {
	vm := subbox.New()
	type T struct {
		A int
		B float64 `json:"b"`
		D func(result any) any
	}
	myT := &T{}
	vm.Set("T", myT)
	vm.Run("T.c = 1;T.a=2;T.b=3.5;T.d = function(data){return data+1}")
	vmT := vm.Get("T").Export().(*T)
	t.Log(vmT == myT)
	t.Log(myT.A, vmT.A)
	t.Log(myT.B, vmT.B)
	t.Log(vm.Run("T.b"))
	t.Log(vm.Run("T.c"))
	t.Log(myT.D(vmT.B))
	t.Log(vm.Run("T.d(T.b)"))
	t.Error("")
}

func TestImportJS(t *testing.T) {
	vm := subbox.New()
	val, err := vm.Import("./d3-array.js")
	t.Log(val, err)
	val, err = vm.Import("./d3-contour.js")
	t.Log(vm.Run("Object.keys(d3)"))

	t.Error(err)
}
