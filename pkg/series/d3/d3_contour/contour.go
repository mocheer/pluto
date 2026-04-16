package d3_contour

import (
	"github.com/samber/lo"
)

// @see https://github1s.com/d3/d3-contour/blob/HEAD/src/contours.js

var cases = [][][][2]float64{
	{},
	{{{1.0, 1.5}, {0.5, 1.0}}},
	{{{1.5, 1.0}, {1.0, 1.5}}},
	{{{1.5, 1.0}, {0.5, 1.0}}},
	{{{1.0, 0.5}, {1.5, 1.0}}},
	{{{1.0, 1.5}, {0.5, 1.0}}, {{1.0, 0.5}, {1.5, 1.0}}},
	{{{1.0, 0.5}, {1.0, 1.5}}},
	{{{1.0, 0.5}, {0.5, 1.0}}},
	{{{0.5, 1.0}, {1.0, 0.5}}},
	{{{1.0, 1.5}, {1.0, 0.5}}},
	{{{0.5, 1.0}, {1.0, 0.5}}, {{1.5, 1.0}, {1.0, 1.5}}},
	{{{1.5, 1.0}, {1.0, 0.5}}},
	{{{0.5, 1.0}, {1.5, 1.0}}},
	{{{1.0, 1.5}, {1.5, 1.0}}},
	{{{0.5, 1.0}, {1.0, 1.5}}},
	{},
}

func Contour() *contour {
	return &contour{
		dx:     1,
		dy:     1,
		smooth: smoothLinear,
		// threshold: d3_array.ThresholdSturges,
	}
}

type contour struct {
	dx        int
	dy        int
	threshold func(values []float64) []float64
	smooth    func(dx, dy int, ring [][2]float64, values []float64, value float64)
}

type fragment struct {
	start float64
	end   float64
	ring  [][2]float64
}

type ContourPolygon struct {
	Type        string           `json:"type"`
	Value       float64          `json:"value"`
	Coordinates [][][][2]float64 `json:"coordinates"`
}

func (m *contour) Contours(values []float64) []*ContourPolygon {
	var tz = m.threshold(values)
	return lo.Map(tz, func(value float64, _ int) *ContourPolygon {
		return m.Contour(values, value)
	})
}

func (m *contour) Contour(values []float64, value float64) *ContourPolygon {
	var polygons = [][][][2]float64{}
	var holes = [][][2]float64{}

	m.isorings(values, value, func(ring [][2]float64) {
		if m.smooth != nil {
			m.smooth(m.dx, m.dy, ring, values, value)
		}
		if Area(ring) > 0 {
			polygons = append(polygons, [][][2]float64{ring})
		} else {
			holes = append(holes, ring)
		}
	})

	lo.ForEach(holes, func(hole [][2]float64, _ int) {
		n := len(polygons)
		for i := 0; i < n; i++ {
			polygon := polygons[i]
			if Contains(polygon[0], hole) != -1 {
				polygon = append(polygon, hole)
				polygons[i] = polygon
				break
			}
		}
	})
	return &ContourPolygon{
		Type:        "MultiPolygon",
		Value:       value,
		Coordinates: polygons,
	}
}

func (m *contour) isorings(values []float64, value float64, callback func(ring [][2]float64)) {
	dx := m.dx
	dy := m.dy
	fragmentByStart := make(map[float64]*fragment)
	fragmentByEnd := make(map[float64]*fragment)
	t1 := values[0] >= value
	x := -1
	y := -1
	//
	var t0, t2, t3 bool

	index := func(point [2]float64) float64 {
		return point[0]*2 + point[1]*float64(m.dx+1)*4
	}
	stitch := func(line [][2]float64, _ int) {
		fx := float64(x)
		fy := float64(y)
		start := [2]float64{line[0][0] + fx, line[0][1] + fy}
		end := [2]float64{line[1][0] + fx, line[1][1] + fy}

		startIndex := index(start)
		endIndex := index(end)

		var f, g *fragment
		if fragmentByEnd[startIndex] != nil {
			f = fragmentByEnd[startIndex]
			if fragmentByStart[endIndex] != nil {
				g = fragmentByStart[endIndex]
				fragmentByEnd[f.end] = nil
				fragmentByStart[g.start] = nil
				if f == g {
					f.ring = append(f.ring, end)
					callback(f.ring)
				} else {
					fr := &fragment{start: f.start, end: g.end, ring: append(f.ring, g.ring...)}
					fragmentByStart[f.start] = fr
					fragmentByEnd[g.end] = fr
				}
			} else {
				fragmentByEnd[f.end] = nil
				f.ring = append(f.ring, end)
				f.end = endIndex
				fragmentByEnd[endIndex] = f
			}
		} else if fragmentByStart[endIndex] != nil {
			f = fragmentByStart[endIndex]
			if fragmentByEnd[startIndex] != nil {
				g = fragmentByEnd[startIndex]
				fragmentByStart[f.start] = nil
				fragmentByEnd[g.end] = nil
				if f == g {
					f.ring = append(f.ring, end)
					callback(f.ring)
				} else {
					fr := &fragment{start: g.start, end: f.end, ring: append(g.ring, f.ring...)}
					fragmentByStart[g.start] = fr
					fragmentByEnd[f.end] = fr
				}
			} else {
				fragmentByStart[f.start] = nil
				f.ring = append([][2]float64{start}, f.ring...)

				f.start = startIndex
				fragmentByStart[startIndex] = f
			}
		} else {
			fr := &fragment{start: startIndex, end: endIndex, ring: [][2]float64{start, end}}
			fragmentByStart[startIndex] = fr
			fragmentByEnd[endIndex] = fr
		}
	}

	lo.ForEach(cases[shiftBool(t1, 1)], stitch)
	for x = 0; x < dx-1; x++ {
		t0 = t1
		t1 = values[x+1] >= value
		lo.ForEach(cases[intBool(t0)|shiftBool(t1, 1)], stitch)
	}

	lo.ForEach(cases[shiftBool(t1, 0)], stitch)

	for y = 0; y < dy-1; y++ {
		x = -1
		t1 = values[y*dx+dx] >= value
		t2 = values[y*dx] >= value
		lo.ForEach(cases[shiftBool(t1, 1)|shiftBool(t2, 2)], stitch)

		for x = 0; x < dx-1; x++ {
			t0 = t1
			t1 = values[y*dx+dx+x+1] >= value
			t3 = t2
			t2 = values[y*dx+x+1] >= value
			lo.ForEach(cases[intBool(t0)|shiftBool(t1, 1)|shiftBool(t2, 2)|shiftBool(t3, 3)], stitch)
		}
		lo.ForEach(cases[intBool(t1)|shiftBool(t2, 3)], stitch)
	}

	x = -1
	t2 = values[y*dx] >= value
	lo.ForEach(cases[shiftBool(t2, 2)], stitch)

	for x = 0; x < dx-1; x++ {
		t3 = t2
		t2 = values[y*dx+x+1] >= value
		lo.ForEach(cases[shiftBool(t2, 2)|shiftBool(t3, 3)], stitch)
	}

	lo.ForEach(cases[shiftBool(t2, 3)], stitch)

}

func (m *contour) Smooth(value func(dx, dy int, ring [][2]float64, values []float64, value float64)) *contour {
	m.smooth = value
	return m
}

func (m *contour) Size(size []int) *contour {
	m.dx = size[0]
	m.dy = size[1]
	return m
}

func (m *contour) Thresholds(legends []float64) *contour {
	m.threshold = func(values []float64) []float64 { return legends }
	return m
}
