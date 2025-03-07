package point_in_polygon_hao

import (
	"testing"

	"github.com/mocheer/xena/pkg/gm"
)

func Test(t *testing.T) {
	polygon := gm.Polygon{
		{
			{1, 1},
			{1, 2},
			{2, 2},
			{2, 1},
			{1, 1},
		},
	}

	t.Log(PointInPolygon(gm.Point{1.5, 1.5}, polygon))
	t.Log(PointInPolygon(gm.Point{4.9, 1.2}, polygon))
	t.Log(PointInPolygon(gm.Point{1, 2}, polygon))
}
