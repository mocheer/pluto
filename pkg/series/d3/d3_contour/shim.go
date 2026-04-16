package d3_contour

// js 位移运算符
// js 支持bool

func shiftBool(value bool, n int) int {
	if value {
		return 1 << n
	}
	return 0
}

func intBool(value bool) int {
	if value {
		return 1
	}
	return 0
}
