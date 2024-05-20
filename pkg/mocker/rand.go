package mocker

import (
	"fmt"
	"math/rand/v2"

	"github.com/mocheer/pluto/pkg/fn"
)

// Bool 生成随机bool值
func Bool() bool {

	return rand.IntN(2) == 0
}

// Caps 生成随机大写字母
func Caps() string {
	return fmt.Sprint(rand.IntN(26) + 65)
}

// LowerCase 生成随机小写字母
func LowerCase() string {
	return fmt.Sprint(rand.IntN(26) + 97)
}

// StringBytes 生成随机字符串
// 可能需要用新的Source来确保rand.IntN的随机性
func StringBytes(num int) []byte {
	result := make([]byte, num)
	for i := 0; i < num; i++ {
		result[i] = byte(rand.IntN(26) + 65)
	}
	return result
}

// String 生成随机字符串
func String(num int) string {
	return fn.BytesToString(StringBytes(num))
}

// RandString 生成随机字符串
func Intn(num int) int {
	return rand.IntN(num)
}
