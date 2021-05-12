package pluto_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/mocheer/pluto/fn"
	"github.com/mocheer/pluto/geom"
	"github.com/mocheer/pluto/rd"
)

// TestFmtString fn.FmtString
func TestFmtString(t *testing.T) {
	result := fn.FmtString(`{a}bcd{e}fg{h}`, map[string]interface{}{"a": "1", "b": 2, "c": 3.0, "h": "4.0"})
	if result != "1bcdfg4.0" {
		t.Error(result)
	}
}

// TestRdString 测试随机字符串
func TestRdString(t *testing.T) {
	result := rd.String(6)
	t.Log(result)
}

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
