package jsg

import (
	"github.com/mocheer/pluto/pkg/ds/ds_text"
	"github.com/robertkrimen/otto"
)

type vm struct {
	Otto *otto.Otto
}

func New() *vm {
	return &vm{Otto: otto.New()}
}

// 导入脚本
func (m vm) Import(fileName string) (otto.Value, error) {
	script, err := ds_text.Read(fileName)
	if err != nil {
		return otto.Value{}, err
	}
	return m.Otto.Run(script)
}
