package distance

import (
	"testing"

	"github.com/mocheer/xena/pkg/gm"
)

// var from = turf.point([-75.343, 39.984]);
// var to = turf.point([-75.534, 39.123]);
// var options = { units: "miles" };

// turf.distance(from, to, options) // = 60.35329997171415

func TestDistance(t *testing.T) {
	d := DistanceLonLat(gm.LonLat{-75.343, 39.984}, gm.LonLat{-75.534, 39.123}, "miles")
	t.Log(d)
}
