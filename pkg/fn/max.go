package fn

import "golang.org/x/exp/constraints"

// Max go1.21提供的max不支持slices切片作为参数
// 不能用comparable 和 std.Number
func Max[T constraints.Ordered](data []T) T {
	max_num := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > max_num {
			max_num = data[i]
		}
	}
	return max_num
}
