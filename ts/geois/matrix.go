package geois

// Matrix 转换矩阵
type Matrix struct {
	Ax float64
	Bx float64
	Cx float64
	Ay float64
	By float64
	Cy float64
}

//Transform 对点进行转换
/**
 * @param	point	待转换的点
 * @return		转换后的点
 */
func (m *Matrix) Transform(point []float64) []float64 {
	var x = point[0]
	var y = point[1]
	return []float64{m.Ax*x + m.Bx*y + m.Cx, m.Ay*x + m.By*y + m.Cy}
}

//UnTransform 对点进行反矩阵转换 即：p = untransform(transform(p))
/**
 * @param	point	待处理的点对象
 * @return		反转换后的对象
 */
func (m *Matrix) UnTransform(point []float64) []float64 {
	var x = point[0]
	var y = point[1]
	return []float64{(x*m.By - y*m.Bx - m.Cx*m.By + m.Cy*m.Bx) / (m.Ax*m.By - m.Ay*m.Bx), (x*m.Ay - y*m.Ax - m.Cx*m.Ay + m.Cy*m.Ax) / (m.Bx*m.Ay - m.By*m.Ax)}
}
