package asy

import "sync"

type Asy struct {
	wg *sync.WaitGroup
}

func New() *Asy {
	return &Asy{wg: &sync.WaitGroup{}}
}

func (m *Asy) Wait() {
	m.wg.Wait()
}

// Add
func (m *Asy) Add(fn func(args ...any), args ...any) *Asy {
	m.wg.Add(1)
	go func() {
		fn(args...)
		m.wg.Done()
	}()
	return m
}

// AddFns
func (m *Asy) AddFns(fns []func()) *Asy {
	run := func(args ...any) {
		for _, fn := range fns {
			fn()
		}
	}
	return m.Add(run)
}

// AddByStep
// @example AddByStep(tasks,16); // 将tasks拆分成16组，并发执行
// 这个分组机制有问题，一般是超海量的程序或者数据请求才需要分组并发执行，但fns容易承受不了
func (m *Asy) AddByStep(fns []func(), step int) *Asy {
	count := len(fns)
	num := count/step + 1
	//
	for i := 0; i < num; i++ {
		start := step * i
		if start < count {
			end := step * (i + 1)
			if end > count {
				end = count
			}
			m.AddFns(fns[start:end])
		} else {
			break
		}
	}
	return m
}

func (m *Asy) AddByNum(fns []func(), num int) *Asy {
	count := len(fns)
	step := count / num
	if step == 0 {
		step = 1
	}
	return m.AddByStep(fns, step)
}
