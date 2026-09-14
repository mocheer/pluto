package clock

import "time"

// SetInterval 间隔 delay 执行一次
func SetInterval(callback func(), delay time.Duration, immediately bool) func() {
	ticker := time.NewTicker(delay)
	go func() {
		for range ticker.C {
			callback()
		}
	}()
	if immediately {
		callback()
	}
	return func() {
		ticker.Stop()
	}
}
