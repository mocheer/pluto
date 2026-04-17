package interpolate

import (
	"math"

	"github.com/mocheer/pluto/pkg/series/turf/distance"
	"github.com/mocheer/pluto/pkg/series/turf/point_grid"
	"github.com/mocheer/pluto/pkg/series/turf/point_in_polygon_hao"
	"github.com/mocheer/xena/pkg/gm"
)

// TODO 如果没有设置bbox，理应从points中获取到bbox

type InterpolatePointOptionsArgs struct {
	CellSize    float64
	GridType    string
	Units       string
	Weight      float64
	BBox        *gm.BBox
	MaskPolygon gm.MultiPolygon
	DistAlg     string
}

type Grid struct {
	Data       []float64
	CellWidth  float64
	CellHeight float64
	Columns    int
	Rows       int
	Xlim       [2]float64
	Ylim       [2]float64
	Zlim       [2]float64
}

func InterpolatePoint(points []gm.PointZ, options InterpolatePointOptionsArgs) (*Grid, error) {
	if options.BBox == nil {
		bbox := gm.NewBBox()
		bbox.ExtendPointZs(points)
		bbox.ExtendBySizeScale(0.1)
		options.BBox = bbox
	}
	// 这里生成的网格需要经纬度坐标，方便用于距离计算
	grid, err := point_grid.NewWithOptions(point_grid.PointGridOptions{
		BBox:     *options.BBox,
		CellSize: options.CellSize,
		Units:    options.Units,
	})
	if err != nil {
		return nil, err
	}
	results := make([]float64, 0, len(grid.Points))
	noDataVal := math.MaxFloat64
	zMin := math.MaxFloat64
	zMax := math.Inf(-1)
	// //
	// const DegreesFactor2 = conversions.DegreesFactor * 2
	// cartographicMap := map[gm.PointZ]gm.Cartographic{}
	// cartographicLatCosMap := map[gm.Cartographic]float64{}
	// getCartographic := func(p gm.PointZ) gm.Cartographic {
	// 	v, ok := cartographicMap[p]
	// 	if !ok {
	// 		v = p.LonLat().ToCartographic()
	// 		cartographicMap[p] = v
	// 		cartographicLatCosMap[v] = math.Cos(v[1])
	// 	}
	// 	return v
	// }
	// distance := func(from gm.Cartographic, p gm.PointZ) float64 {
	// 	to := getCartographic(p)
	// 	//
	// 	dLatRad := to[1] - from[1]
	// 	dLonRad := to[0] - from[0]
	// 	lat1Rad := from[1]
	// 	lat2RadCos := cartographicLatCosMap[to]
	// 	// Haversine 公式
	// 	// 这是一个用于计算两个经度和纬度之间的距离的公式。
	// 	dLatRad2 := math.Sin(dLatRad / 2)
	// 	dLonRad2 := math.Sin(dLonRad / 2)
	// 	a := dLatRad2*dLatRad2 + dLonRad2*dLonRad2*math.Cos(lat1Rad)*lat2RadCos
	// 	return DegreesFactor2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	// }
	// var distFunc func(from, to gm.LonLat) float64
	for _, grid := range grid.Points {
		var zw, sw float64
		show := true
		if options.MaskPolygon != nil {
			isIn, _ := point_in_polygon_hao.PointInMultiPolygon(grid.ToPoint(), options.MaskPolygon)
			show = isIn
		}
		zVal := noDataVal
		if show {
			for _, p := range points {
				// d := distance.DistanceCartographic(grid.ToCartographic(), p.LonLat().ToCartographic())
				// d := distance.DistanceCartographicSimple(grid, p.LonLat().ToCartographic())
				d := distance.DistanceLonLatSimple(grid, p.ToLonLat())
				zValue := p[2]
				if d == 0 { //当前格点刚好是一个测站的位置
					zw = zValue
				}
				w := 1.0 / math.Pow(d, options.Weight)
				sw += w
				zw += w * zValue
			}
			zVal = zw / sw
		}
		if zMin > zVal {
			zMin = zVal
		}
		if zMax < zVal {
			zMax = zVal
		}
		results = append(results, zVal)
	}
	resultGrid := &Grid{
		Data:       results,
		Rows:       grid.Rows,
		Columns:    grid.Columns,
		Xlim:       grid.Xlim,
		Ylim:       grid.Ylim,
		Zlim:       [2]float64{zMin, zMax},
		CellWidth:  grid.CellWidth,
		CellHeight: grid.CellHeight,
	}
	return resultGrid, nil
}

// InterpolatePoint2
// turf 的Interpolate算法，本身就是生成FeatureCollection，这里简化为[]gm.PointZ
// func InterpolatePointWithTurf(points []gm.PointZ, options InterpolatePointOptionsArgs) ([]gm.PointZ, error) {
// 	if options.BBox == nil {
// 		bbox := gm.NewBBox()
// 		bbox.ExtendPointZs(points)
// 		options.BBox = bbox
// 	}
// 	// 这里生成的网格需要经纬度坐标，方便用于距离计算
// 	grid, err := point_grid.NewWithOptions(point_grid.PointGridOptions{
// 		BBox:     *options.BBox,
// 		CellSize: options.CellSize,
// 		Units:    options.Units,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	results := make([]gm.PointZ, 0, len(grid.Points))
// 	zMin := math.Inf(+1)
// 	zMax := math.Inf(-1)
// 	for _, grid := range grid.Points {
// 		var zw, sw float64
// 		for _, p := range points {
// 			d := distance.Distance(grid, p.LonLat(), options.Units)
// 			zValue := p[2]
// 			if d == 0 {
// 				zw = zValue
// 			}
// 			w := 1.0 / math.Pow(d, options.Weight)
// 			sw += w
// 			zw += w * zValue
// 		}
// 		np := grid.ToPointZ()
// 		np[2] = zw / sw
// 		//
// 		if options.Mask != nil {
// 			isIn, _ := point_in_polygon_hao.PointInMultiPolygon(np.Point(), options.Mask)
// 			if !isIn {
// 				np[2] = 9999.0
// 			}
// 		}
// 		if zMin > np[2] {
// 			zMin = np[2]
// 		}
// 		if zMax < np[2] {
// 			zMax = np[2]
// 		}
// 		results = append(results, np)
// 	}

// 	return results, nil
// }
