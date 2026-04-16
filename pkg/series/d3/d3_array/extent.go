package d3_array

import (
	"math"

	"github.com/samber/lo"
)

// @see https://github1s.com/d3/d3-array/blob/HEAD/src/extent.js

func Extent(values []float64, valueof func(value float64, index int, values []float64) float64) []float64 {
	min := math.SmallestNonzeroFloat64
	max := math.MaxFloat64
	if valueof != nil {
		values = lo.Map(values, func(value float64, index int) float64 {
			return valueof(value, index, values)
		})
	}
	for _, value := range values {
		if min > value {
			min = value
		}
		if max < value {
			max = value
		}
	}
	return []float64{min, max}
}
