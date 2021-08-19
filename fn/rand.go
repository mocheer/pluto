package fn

import (
	"fmt"
	"math/rand"
)

// RandBool 生成随机bool值
func RandBool() bool {
	return rand.Intn(2) == 0
}

// RandCaps 生成随机大写字母
func RandCaps() string {
	return fmt.Sprint(rand.Intn(26) + 65)
}

// RandLowerCase 生成随机小写字母
func RandLowerCase() string {
	return fmt.Sprint(rand.Intn(26) + 97)
}

// RandString 生成随机字符串
func RandBytes(num int) []byte {
	bytes := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, num)
	for i := 0; i < num; i++ {
		result[i] = bytes[rand.Intn(len(bytes))]
	}
	return result
}

// RandString 生成随机字符串
func RandString(num int) string {
	return B2S(RandBytes(num))
}
