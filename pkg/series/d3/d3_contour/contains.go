package d3_contour

func Contains(ring [][2]float64, hole [][2]float64) int {
	n := len(hole)
	for i := 0; i < n; i++ {
		c := RingContains(ring, hole[i])
		if c != 0 {
			return c
		}
	}
	return 0
}

func RingContains(ring [][2]float64, point [2]float64) int {
	x := point[0]
	y := point[1]
	contains := -1
	n := len(ring)
	j := n - 1
	for i := 0; i < n; i++ {
		pi := ring[i]
		xi := pi[0]
		yi := pi[1]
		pj := ring[j]
		xj := pj[0]
		yj := pj[1]
		if SegmentContains(pi, pj, point) {
			return 0
		}
		if ((yi > y) != (yj > y)) && (x < (xj-xi)*(y-yi)/(yj-yi)+xi) {
			contains = -contains
		}

		j = i
	}
	return contains
}

func SegmentContains(a, b, c [2]float64) bool {
	i := 0
	if a[0] == b[0] {
		i = 1
	}
	return Collinear(a, b, c) && Within(a[i], c[i], b[i])
}

func Collinear(a, b, c [2]float64) bool {
	return (b[0]-a[0])*(c[1]-a[1]) == (c[0]-a[0])*(b[1]-a[1])
}

func Within(p, q, r float64) bool {
	return p <= q && q <= r || r <= q && q <= p
}
