package d3_array

import (
	"math"
)

// https://github1s.com/d3/d3-array/blob/HEAD/src/range.js

func Range(start, stop, step float64) []float64 {
	n := int(math.Max(0.0, math.Ceil((stop-start)/step))) | 0
	result := make([]float64, n)
	for i := 0; i < n; i++ {
		result[i] = start + float64(i)*step
	}

	return result
}
