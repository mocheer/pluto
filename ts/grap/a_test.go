package grap_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/mocheer/pluto/ts/grap"
)

// TestBezierCurve 测试贝塞尔曲线
func TestBezierCurve(t *testing.T) {
	n := 5
	var data []grap.Point
	for index := 0; index < n; index++ {
		data = append(data, grap.Point{X: rand.Float64() * 800, Y: rand.Float64() * 500})
	}
	bezierCurve := grap.NewBezierCurve(data)
	points := bezierCurve.GetPoints(0.01)
	for index, point := range points {
		fmt.Println(index, point.X, point.Y)
	}
}
