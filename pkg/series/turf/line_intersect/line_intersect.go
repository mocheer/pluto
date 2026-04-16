package line_intersect

import (
	"fmt"
	"math"

	"github.com/mocheer/xena/pkg/gm"
)

const (
	tolerance = 1e-12
)

// 计算线段交点，返回所有交点数据
func LineIntersect(g1, g2 gm.LineString) []gm.Point {
	segments1 := lineToSegments(g1)
	segments2 := lineToSegments(g2)

	seen := make(map[string]bool)
	features := []gm.Point{}

	for _, seg1 := range segments1 {
		for _, seg2 := range segments2 {
			if point, ok := segmentIntersection(seg1, seg2); ok {
				key := fmt.Sprintf("%.6f,%.6f", point[0], point[1])
				if !seen[key] {
					seen[key] = true
					features = append(features, point)
				}
			}
		}
	}

	return features
}

// 将LineString拆分为线段
func lineToSegments(points gm.LineString) []gm.Line {
	segments := make([]gm.Line, len(points)-1)
	for i := 0; i < len(points)-1; i++ {
		segments[i] = gm.Line{points[i], points[i+1]}
	}
	return segments
}

// 计算两条线段的交点
func segmentIntersection(seg1, seg2 gm.Line) (gm.Point, bool) {
	a, b := seg1[0], seg1[1]
	c, d := seg2[0], seg2[1]

	denom := (d[1]-c[1])*(b[0]-a[0]) - (d[0]-c[0])*(b[1]-a[1])

	// 处理平行或重合的情况
	if math.Abs(denom) < tolerance {
		return gm.Point{}, false
	}

	numeA := (d[0]-c[0])*(a[1]-c[1]) - (d[1]-c[1])*(a[0]-c[0])
	numeB := (b[0]-a[0])*(a[1]-c[1]) - (b[1]-a[1])*(a[0]-c[0])

	ua := numeA / denom
	ub := numeB / denom

	// 检查交点是否在线段上
	if ua < -tolerance || ua > 1+tolerance || ub < -tolerance || ub > 1+tolerance {
		return gm.Point{}, false
	}

	// 计算实际交点坐标
	x := a[0] + ua*(b[0]-a[0])
	y := a[1] + ua*(b[1]-a[1])

	return gm.Point{x, y}, true
}
