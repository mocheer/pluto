package ctp

import "sync"

type CtpPool struct {
	funcs []func()
	num   int
	wg    *sync.WaitGroup
}

func NewPool() *CtpPool {
	return &CtpPool{
		num: 5,
		wg:  &sync.WaitGroup{},
	}
}

func (m *CtpPool) Save(uri string, fileName string) {
	m.funcs = append(m.funcs, func() {
		New().Save(uri, fileName)
		m.wg.Done()
	})
	m.wg.Add(1)
	if len(m.funcs) >= m.num {
		m.Wait()
	}
}

func (m *CtpPool) Wait() {
	for _, fn := range m.funcs {
		go fn()
	}
	m.wg.Wait()
	m.funcs = []func(){}

}
