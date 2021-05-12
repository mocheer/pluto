package geom

//Point
type Point struct {
	X float64
	Y float64
}

//Line
type Line struct {
	Start Point
	End   Point
}

//Lines
type Lines struct {
	Data []Point
}

//BezierCurve interface
type BezierCurve interface {
	GetPoints(step float64) []Point
	GetPoint(t float64) Point
}
