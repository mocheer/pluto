package fn

import "sync"

// GoFnsGroupBy
func GoFnsGroupBy(count int, num int, fns []func()) *sync.WaitGroup {
	size := len(fns)
	if size > count {
		start := 0
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			for start < size {
				end := start + count
				if end > size {
					end = size
				}
				GoFns(num, fns[start:end]).Wait()
				start += count
			}
			wg.Done()
		}()
		return wg
	}
	return GoFns(num, fns)
}
