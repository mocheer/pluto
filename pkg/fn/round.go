package fn

import "math"

// Round 四舍五入保留小数
func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

// RoundInt 四舍五入取整数
func RoundInt(f float64) int {
	return int(f + 0.5)
}
