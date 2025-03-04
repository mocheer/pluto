package fn

import "golang.org/x/exp/constraints"

// Min go1.21提供的min不支持slices切片作为参数
// 不能用comparable 和 std.Number
func Min[T constraints.Ordered](data []T) T {
	min_num := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] < min_num {
			min_num = data[i]
		}
	}
	return min_num
}
