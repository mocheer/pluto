package rd

import (
	"fmt"
	"math/rand"
)

// Bool 随机bool值
func Bool() bool {
	return rand.Intn(2) == 0
}

// Caps 随机大写字母
func Caps() string {
	return fmt.Sprint(rand.Intn(26) + 65)
}

// LowerCase 随机小写字母
func LowerCase() string {
	return fmt.Sprint(rand.Intn(26) + 97)
}

// String 生成随机字符串
func String(num int) string {
	bytes := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, num)
	for i := 0; i < num; i++ {
		result[i] = bytes[rand.Intn(len(bytes))]
	}
	return string(result)
}
