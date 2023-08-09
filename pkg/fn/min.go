package fn

import "github.com/mocheer/pluto/pkg/itypes"

// Min go1.21提供的min不支持slices切片作为参数
func Min[T itypes.Number](data []T) T {
	min_num := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] < min_num {
			min_num = data[i]
		}
	}

	return min_num
}
