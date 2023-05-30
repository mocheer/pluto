package fn

import "github.com/mocheer/pluto/pkg/itypes"

func Min[T itypes.Number](data []T) T {
	min_num := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] < min_num {
			min_num = data[i]
		}
	}
	return min_num
}
