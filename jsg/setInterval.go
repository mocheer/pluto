package jsg

import "time"

// SetInterval 间隔delay毫秒执行一次
func SetInterval(callback func(), delay time.Duration) func() {
	ticker := time.NewTicker(time.Millisecond * delay)
	go func() {
		for range ticker.C {
			callback()
		}
	}()
	return func() {
		ticker.Stop()
	}
}
