package fn

// UnFlat2 将一个一维切片转成二维切片
func UnFlat2(data []int, numCols int) [][]int {
	if numCols == 0 {
		return nil
	}
	numRows := (len(data) + numCols - 1) / numCols
	result := make([][]int, numRows)
	for i := 0; i < len(data); i += numCols {
		end := i + numCols
		if end > len(data) {
			end = len(data)
		}
		row := i / numCols
		result[row] = data[i:end]
	}

	return result
}
