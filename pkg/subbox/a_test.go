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
	t.Log(val.Object().Get("a"))
	t.Log(vm.Otto.Run("JSON.stringify(module)"))
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
