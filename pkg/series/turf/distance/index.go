package distance

import (
	"math"

	"github.com/mocheer/pluto/pkg/turf/conversions"
	"github.com/mocheer/xena/pkg/gm"
)

// 这个变量很多地方都有，后面需要整理下
const RADIANS_PER_DEGREE = math.Pi / 180 // 每一个角度单位对应的弧度值

// DistanceLonLat 计算两个地理坐标点之间的距离
// 参数：
//
//	from: 起点坐标
//	to: 终点坐标
//	units: 距离单位（默认为千米）
//
// 返回：
//
//	距离值
func DistanceLonLat(from, to gm.LonLat, units string) float64 {
	// 将纬度和经度从度转换为弧度
	dLat := to.Lat() - from.Lat()
	dLon := to.Lon() - from.Lon()
	lat1 := from.Lat()
	lat2 := to.Lat()

	dLatRad := dLat * RADIANS_PER_DEGREE
	dLonRad := dLon * RADIANS_PER_DEGREE
	lat1Rad := lat1 * RADIANS_PER_DEGREE
	lat2Rad := lat2 * RADIANS_PER_DEGREE

	// Haversine 公式
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

const DegreesFactor2 = conversions.DegreesFactor * 2

// 哈弗辛公式
func DistanceCartographic(from, to gm.Cartographic) float64 {
	//
	lat1Rad := from[1]
	lat2Rad := to[1]
	dLatRad := lat2Rad - lat1Rad
	dLonRad := to[0] - from[0]
	// Haversine 公式
	// 这是一个用于计算两个经度和纬度之间的距离的公式。
	dLatRad2 := math.Sin(dLatRad / 2)
	dLonRad2 := math.Sin(dLonRad / 2)
	a := dLatRad2*dLatRad2 + dLonRad2*dLonRad2*math.Cos(lat1Rad)*math.Cos(lat2Rad)
	return DegreesFactor2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}

// DistanceLonLatSimple
// 无投影平面近似法，误差较大
func DistanceLonLatSimple(from, to gm.LonLat) float64 {
	//
	dLatRad := to[1] - from[1]
	dLonRad := to[0] - from[0]
	return math.Sqrt(dLonRad*dLonRad+dLatRad*dLatRad) * 111000
}
