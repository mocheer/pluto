package geom_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/mocheer/pluto/ts/geom"
)

// TestBezierCurve 测试贝塞尔曲线
func TestBezierCurve(t *testing.T) {
	n := 5
	var data []geom.Point
	for index := 0; index < n; index++ {
		data = append(data, geom.Point{X: rand.Float64() * 800, Y: rand.Float64() * 500})
	}
	bezierCurve := geom.NewBezierCurve(data)
	points := bezierCurve.GetPoints(0.01)
	for index, point := range points {
		fmt.Println(index, point.X, point.Y)
	}
}
