package clock

import (
	"fmt"
	"time"
)

// Cost
// defer Cost()("执行时间耗时")
func Cost() func(name string) {
	start := time.Now()
	return func(name string) {
		tc := time.Since(start)
		fmt.Printf("%v :  %v s\n", name, tc.Seconds())
	}
}
