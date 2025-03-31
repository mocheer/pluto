package point_in_polygon_hao

import (
	"errors"

	"github.com/mocheer/xena/pkg/gm"
)

// PointInMultiPolygon
func PointInMultiPolygon(p gm.Point, polygons gm.MultiPolygon) (bool, error) {
	for _, polygon := range polygons {
		isIn, _ := PointInPolygon(p, polygon)
		if isIn {
			return true, nil
		}
	}
	return false, nil
}

// PointInPolygon checks if a point is inside a polygon.
func PointInPolygon(p gm.Point, polygon gm.Polygon) (bool, error) {
	k := 0

	for _, contour := range polygon {
		if len(contour) < 4 {
			return false, errors.New("contour must have at least 4 points (including the closing point)")
		}

		if contour[0] != contour[len(contour)-1] {
			return false, errors.New("first and last coordinates in a ring must be the same")
		}

		currentP := contour[0]
		u1 := currentP[0] - p[0]
		v1 := currentP[1] - p[1]
		// TODO 检测对比
		// https://github.com/rowanwins/point-in-polygon-hao/blob/master/src/index.js
		for ii := 0; ii < len(contour); ii++ {
			nextP := contour[ii]

			u2 := nextP[0] - p[0]
			v2 := nextP[1] - p[1]

			if v1 == 0 && v2 == 0 {
				if (u2 <= 0 && u1 >= 0) || (u1 <= 0 && u2 >= 0) {
					return false, nil
				}
			} else if (v2 >= 0 && v1 <= 0) || (v2 <= 0 && v1 >= 0) {
				f := Orient2d(u1, v1, u2, v2, 0, 0)
				if f == 0 { //这里要根据配置返回true或者false，有些判断要支持边界
					return false, nil
				}
				if (f > 0 && v2 > 0 && v1 <= 0) || (f < 0 && v2 <= 0 && v1 > 0) {
					k++
				}
			}

			currentP = nextP
			v1 = v2
			u1 = u2
		}
	}

	return k%2 != 0, nil
}
