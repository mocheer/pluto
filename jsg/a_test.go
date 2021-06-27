package jsg_test

import (
	"testing"

	"github.com/mocheer/pluto/jsg"
)

func TestRun(t *testing.T) {
	vm := jsg.New()
	val, err := vm.Otto.Run("var module ={};module.exports = {a:1};")
	if err != nil {
		t.Error(err)
	}
	t.Log(val.Object().Get("a"))
	t.Log(vm.Otto.Run("JSON.stringify(module)"))
}
