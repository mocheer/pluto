package boolean_overlap

import (
	"github.com/mocheer/pluto/pkg/turf/line_intersect"
	"github.com/mocheer/xena/pkg/gm"
)

// BooleanOverlap  判断多边形和多边形是否相交重叠
func BooleanOverlapPolygon(poly1, poly2 gm.Polygon) bool {
	for _, line1 := range poly1 {
		for _, line2 := range poly2 {
			points := line_intersect.LineIntersect(line1, line2)
			if len(points) > 0 {
				return true
			}
		}
	}
	return false
}
