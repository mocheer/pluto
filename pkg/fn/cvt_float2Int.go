package fn

import "math"

func Float64ToInt16(v float64) int16 {
	if math.IsNaN(v) {
		return 0
	}
	if v > math.MaxInt16 {
		return math.MaxInt16
	} else if v < math.MinInt16 {
		return math.MinInt16
	}
	return int16(math.Round(v))
}
