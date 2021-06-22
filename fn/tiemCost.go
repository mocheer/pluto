package fn

import (
	"fmt"
	"time"
)

// TimeCost
// defer TimeCost()("执行时间耗时")
func TimeCost() func(name string) {
	start := time.Now()
	return func(name string) {
		tc := time.Since(start)
		fmt.Printf("%v :  %v s\n", name, tc.Seconds())
	}
}
