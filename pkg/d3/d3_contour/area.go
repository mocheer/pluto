package d3_contour

func area(ring [][2]float64) float64 {
	n := len(ring)
	result := ring[n-1][1]*ring[0][0] - ring[n-1][0]*ring[0][1]
	for i := 1; i < n; i++ {
		result += ring[i-1][1]*ring[i][0] - ring[i-1][0]*ring[i][1]
	}
	return result
}
