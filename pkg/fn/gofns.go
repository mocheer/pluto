package fn

import (
	"sync"
)

// GoFns
// @example GoFns(16,tasks); // 将tasks拆分成16组，并发执行
// 这个分组机制有问题
func GoFns(step int, fns []func()) (wg *sync.WaitGroup) {
	size := len(fns)
	num := size/step + 1
	wg = &sync.WaitGroup{}
	for i := 0; i < num; i++ {
		start := step * i
		if start < size {
			end := step * (i + 1)
			if end > size {
				end = size
			}
			wg.Add(1) //添加一个计数
			go func(fns []func()) {
				for _, fn := range fns {
					fn()
				}
				wg.Done() //减去一个计数
			}(fns[start:end])
		} else {
			break
		}
	}
	return
}
