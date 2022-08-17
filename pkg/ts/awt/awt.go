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

func (m *Awt) Wait() *Awt {
	if len(m.source) > 0 {
		m.Asy.AddFns(m.source)
		m.source = []func(){}

	}

	m.Asy.Wait()
	m.num = 0
	return m
}
