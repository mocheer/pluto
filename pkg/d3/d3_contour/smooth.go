package d3_contour

func smoothLinear(dx, dy int, ring [][2]float64, values []float64, value float64) {
	for i := range ring {
		var x = ring[i][0]
		var y = ring[i][1]
		var xt = int(x)
		var yt = int(y)
		var v1_index = yt*dx + xt //有可能溢出
		if v1_index < len(values) {
			var v1 = values[v1_index]
			if x > 0 && xt < dx && float64(xt) == x {
				var v0 = values[v1_index-1]
				ring[i][0] = x + (value-v0)/(v1-v0) - 0.5
			}
			if y > 0 && yt < dy && float64(yt) == y {
				var v0 = values[(yt-1)*dx+xt]
				ring[i][1] = y + (value-v0)/(v1-v0) - 0.5
			}
		}
	}
}
