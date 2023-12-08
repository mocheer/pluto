package awt

import (
	"github.com/mocheer/pluto/pkg/ts/asy"
)

type Awt struct {
	Asy         *asy.Asy
	Num         int
	MaxAsyCount int
	//
	num    int
	source []func()
}

// New
// num				 协程个数，一旦超过，就会等待所有协程执行完毕，后续应该改成当一个协程执行完毕，num-=1
// maxAsyCount 每个协程执行的程序个数
func New(num int, maxAsyCount int) *Awt {
	return &Awt{Asy: asy.New(), Num: num, MaxAsyCount: maxAsyCount, num: 0, source: []func(){}}
}

func (m *Awt) Add(fn func(args ...any), args ...any) *Awt {
	m.source = append(m.source, func() {
		fn(args...)
	})
	if len(m.source) >= m.MaxAsyCount {
		m.Asy.AddFns(m.source)
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
		m.Asy.AddFns(m.source)
		m.source = []func(){}

	}

	m.Asy.Wait()
	m.num = 0
	return m
}
