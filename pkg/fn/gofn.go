package fn

import (
	"sync"
)

type GoFnStruct struct {
	Wg *sync.WaitGroup
}

//
func (m GoFnStruct) Go(fn func(args ...interface{}), args ...interface{}) {
	m.Wg.Add(1)
	go func() {
		fn(args...)
		m.Wg.Done()
	}()
}
