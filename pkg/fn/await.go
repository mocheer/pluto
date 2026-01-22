package fn

import "sync"

// Go
// 有点类似Promise.all的用法
func Go(fns ...func()) {
	var wg sync.WaitGroup
	if len(fns) > 0 {
		for _, f := range fns {
			wg.Go(f)
		}
		wg.Wait()
	}
}
