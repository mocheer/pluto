package fn

import (
	"time"
)

// SetSleep 循环执行，同步等待
// 不同于setInterval的非阻塞执行，这里会停止直到返回false
func SetSleep(action func() bool, delay int) {
	for action() {
		time.Sleep(time.Duration(delay) * time.Second)
	}
}
