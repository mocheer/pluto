package d3_array

import "math"

// @see https://github1s.com/d3/d3-array/blob/HEAD/src/threshold/sturges.js#L1-L6

func ThresholdSturges(values []float64) float64 {
	return math.Ceil(math.Log(Count(values, nil))/math.Ln2) + 1
}
