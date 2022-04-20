package mocker

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mocheer/pluto/pkg/fn"
)

// Bool 生成随机bool值
func Bool() bool {
	return rand.Intn(2) == 0
}

// Caps 生成随机大写字母
func Caps() string {
	return fmt.Sprint(rand.Intn(26) + 65)
}

// LowerCase 生成随机小写字母
func LowerCase() string {
	return fmt.Sprint(rand.Intn(26) + 97)
}

// StringBytes 生成随机字符串
func StringBytes(num int) []byte {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, num)
	for i := 0; i < num; i++ {
		result[i] = byte(r.Intn(26) + 65)
	}
	return result
}

// String 生成随机字符串
func String(num int) string {
	return fn.B2S(StringBytes(num))
}

// RandString 生成随机字符串
func Intn(num int) int {
	return rand.Intn(num)
}
