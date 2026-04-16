package d3_array

import (
	"github.com/samber/lo"
)

// @see https://github1s.com/d3/d3-array/blob/HEAD/src/count.js#L1-L19
// d3.count([{n: "Alice", age: NaN}, {n: "Bob", age: 18}, {n: "Other"}], d => d.age) // 1
// d3.count([1, 2, undefined, 3, NaN, 4]) 4
// d3.count([new Date("2019-01-01"), new Date(NaN)]) // 1

// Count
func Count(values []float64, valueof func(value any, index int, values []float64) float64) float64 {
	if valueof != nil {
		values = lo.Map(values, func(value float64, index int) float64 {
			return valueof(value, index, values)
		})
	}
	return float64(len(values))
}
