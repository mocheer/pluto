package awt

import (
	"sync"
)

type Awt struct {
	Asy         *sync.WaitGroup
	Num         int
	MaxAsyCount int
	//
	num    int
	source []func()
}

// New
// num	协程个数，一旦超过，就会等待所有协程执行完毕，后续应该改成当一个协程执行完毕，num-=1
// maxAsyCount 每个协程执行的程序个数
func New(num int, maxAsyCount int) *Awt {
	return &Awt{Asy: &sync.WaitGroup{}, Num: num, MaxAsyCount: maxAsyCount, num: 0, source: []func(){}}
}

// Add
func (m *Awt) Add(fn func(args ...any), args ...any) *Awt {
	m.source = append(m.source, func() {
		fn(args...)
	})
	if len(m.source) >= m.MaxAsyCount {
		for _, fn := range m.source {
			m.Asy.Go(fn)
		}
		m.num++
		if m.num >= m.Num {
			m.Asy.Wait()
			m.num = 0
		}
		m.source = []func(){}
	}
	return m
}

// Wait 防止程序个数小于一个协程的个数，没有添加进协程执行
func (m *Awt) Wait() *Awt {
	if len(m.source) > 0 {
		for _, fn := range m.source {
			m.Asy.Go(fn)
		}
		m.source = []func(){}
	}
	m.Asy.Wait()
	m.num = 0
	return m
}
