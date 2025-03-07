package subbox

import (
	"errors"

	"github.com/dop251/goja"
	"github.com/mocheer/pluto/pkg/ds/ds_text"
)

type vm struct {
	Ctx *goja.Runtime
}

func New() *vm {
	ctx := goja.New()
	// 将结构体的成员变量和方法都变成小写（跟otto一致，否则默认首字母都是大写的）
	ctx.SetFieldNameMapper(goja.UncapFieldNameMapper())
	return &vm{Ctx: ctx}
}

// Import 导入脚本
func (m vm) Import(fileName string) (goja.Value, error) {
	script, err := ds_text.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return m.Ctx.RunString(script)
}

// Get
func (m vm) Get(key string) goja.Value {
	return m.Ctx.Get(key)
}

// Set
func (m vm) Set(key string, value any) error {
	return m.Ctx.Set(key, value)
}

// Run
func (m vm) Run(script string) (goja.Value, error) {
	return m.Ctx.RunString(script)
}

// Global
func (m vm) Global() goja.Value {
	return m.Ctx.GlobalObject()
}

// Call
func (m vm) Call(key string, this goja.Value, args ...goja.Value) (goja.Value, error) {
	fn, ok := goja.AssertFunction(m.Get(key))
	if !ok {
		return nil, errors.New("方法名不存在")
	}
	return fn(this, args...)
}

// Call2
func (m vm) Call2(key string, args ...goja.Value) (goja.Value, error) {
	return m.Call(key, goja.Undefined())
}
