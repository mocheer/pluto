package distance

import (
	"math"

	"github.com/mocheer/pluto/pkg/turf/conversions"
	"github.com/mocheer/xena/pkg/gm"
)

// 这个变量很多地方都有，后面需要整理下
const RADIANS_PER_DEGREE = math.Pi / 180 // 每一个角度单位对应的弧度值

// Distance 计算两个地理坐标点之间的距离
// 参数：
//
//	from: 起点坐标
//	to: 终点坐标
//	units: 距离单位（默认为千米）
//
// 返回：
//
//	距离值
func Distance(from, to gm.LonLat, units string) float64 {
	// 将纬度和经度从度转换为弧度
	dLat := to.Lat() - from.Lat()
	dLon := to.Lon() - from.Lon()
	lat1 := from.Lat()
	lat2 := to.Lat()

	dLatRad := dLat * RADIANS_PER_DEGREE
	dLonRad := dLon * RADIANS_PER_DEGREE
	lat1Rad := lat1 * RADIANS_PER_DEGREE
	lat2Rad := lat2 * RADIANS_PER_DEGREE

	// Haversine公式
	// 这是一个用于计算两个经度和纬度之间的距离的公式。
	a := math.Pow(math.Sin(dLatRad/2), 2) +
		math.Pow(math.Sin(dLonRad/2), 2)*
			math.Cos(lat1Rad)*math.Cos(lat2Rad)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	d, err := conversions.RadiansToLength(c, units)
	if err != nil {
		panic(err)
	}
	return d
}
