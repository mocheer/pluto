package d3_array

import (
	"math"

	"github.com/samber/lo"
)

// https://github1s.com/d3/d3-array/blob/HEAD/src/ticks.js#L5

var e10 = math.Sqrt(50)
var e5 = math.Sqrt(10)
var e2 = math.Sqrt(2)

func Ticks(start, stop, count float64) []float64 {

	if start == stop && count > 0 {
		return []float64{start}
	}
	reverse := stop < start

	var n int
	var step float64
	if reverse {
		n = int(start)
		start = stop
		stop = float64(n)
	}

	step = TickIncrement(start, stop, count)
	if step == 0 {
		return []float64{}
	}
	var ticks []float64
	if step > 0 {
		r0 := math.Round(start / step)
		r1 := math.Round(stop / step)
		if r0*step < start {
			r0++
		}
		if r1*step > stop {
			r1--
		}
		n = int(r1 - r0 + 1)
		ticks = make([]float64, n)

		for i := 0; i < n; i++ {
			ticks[i] = (r0 + float64(i)) * step
		}

	} else {
		step = -step
		r0 := math.Round(start * step)
		r1 := math.Round(stop * step)
		if r0/step < start {
			r0++
		}
		if r1/step > stop {
			r1--
		}
		n = int(r1 - r0 + 1)
		ticks = make([]float64, n)

		for i := 0; i < n; i++ {
			ticks[i] = (r0 + float64(i)) / step
		}
	}

	if reverse {
		ticks = lo.Reverse(ticks)
	}

	return ticks
}

func TickIncrement(start, stop, count float64) float64 {
	step := (stop - start) / math.Max(0, count)
	power := math.Floor(math.Log(step) / math.Ln10)
	err := step / math.Pow(10, power)

	var val float64

	switch {
	case err >= e10:
		val = 10
	case err >= e5:
		val = 5
	case err >= e2:
		val = 2
	default:
		val = 1
	}

	if power >= 0 {
		return val * math.Pow(10, power)
	}

	return -math.Pow(10, -power) / val

}

func TickStep(start, stop, count float64) float64 {
	step0 := math.Abs(stop-start) / math.Max(0, count)
	step1 := math.Pow(10, math.Floor(math.Log(step0)/math.Ln10))
	err := step0 / step1
	if err >= e10 {
		step1 = step1 * 10
	} else if err >= e5 {
		step1 = step1 * 5
	} else if err >= e2 {
		step1 = step1 * 2
	}
	if stop < start {
		return -step1
	}
	return step1
}
