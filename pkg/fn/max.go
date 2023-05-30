package fn

import "github.com/mocheer/pluto/pkg/itypes"

func Max[T itypes.Number](data []T) T {
	max_num := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > max_num {
			max_num = data[i]
		}
	}
	return max_num
}
