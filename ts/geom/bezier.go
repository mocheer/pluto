package geom

import (
	"math"
)

// bezierCurve implements the BezierCurve interface.
type bezierCurve struct {
	data []Point
}

//NewBezierCurve creates a new bezier curve.
func NewBezierCurve(data []Point) BezierCurve {
	return &bezierCurve{data}
}

//GetPoints step [0,1]
func (m *bezierCurve) GetPoints(step float64) []Point {
	var points []Point
	for t := 0.0; t <= 1; t += step {
		points = append(points, m.GetPoint(t))
	}
	return points
}

//GetPoint t [0,1]
func (m *bezierCurve) GetPoint(t float64) Point {
	data := m.data
	var x float64
	var y float64
	n := len(data)
	for i, p := range data {
		b := formula(n-1, i, t)
		x += p.X * b
		y += p.Y * b
	}
	return (Point{x, y})
}

//基函数
func formula(n int, k int, t float64) float64 {
	fk := float64(k)
	fn := float64(n)
	return c(n, k) * math.Pow(t, fk) * math.Pow(1-t, fn-fk)
}

//组合排序
func c(n int, k int) float64 {
	son := factorial(n)
	mother := factorial(k) * factorial(n-k)
	return float64(son) / float64(mother)
}

//阶乘
func factorial(i int) int {
	n := 1
	for j := 1; j <= i; j++ {
		n *= j
	}
	return n
}
