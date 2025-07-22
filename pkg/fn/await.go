package fn

import "sync"

// 目前以下功能已经在规划中了
type WaitGroup struct {
	*sync.WaitGroup
}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{&sync.WaitGroup{}}
}

// var wg sync.WaitGroup
// wg.Go()
// wg.Go()
// wg.Wait()
func (wg *WaitGroup) Go(f func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		f()
	}()
}
