package subbox

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

// Import 导入脚本
func (m vm) Import(fileName string) (otto.Value, error) {
	script, err := ds_text.Read(fileName)
	if err != nil {
		return otto.Value{}, err
	}
	return m.Otto.Run(script)
}

// Get
func (m vm) Get(key string) (otto.Value, error) {
	return m.Otto.Run(key)
}

// Set
func (m vm) Set(key string, value any) error {
	return m.Otto.Set(key, value)
}

// Run
func (m vm) Run(script string) (otto.Value, error) {
	return m.Otto.Run(script)
}
