package fn

import "github.com/mocheer/pluto/pkg/itypes"

// Max go1.21提供的max不支持slices切片作为参数
func Max[T itypes.Number](data []T) T {
	max_num := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > max_num {
			max_num = data[i]
		}
	}
	return max_num
}
