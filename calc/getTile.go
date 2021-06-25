package calc

import (
	"math"

	"github.com/mocheer/pluto/ts/geois"
	"github.com/mocheer/pluto/ts/grap"
)

// LonLat2Tile 经纬度转成瓦片
func LonLat2Tile(lon float64, lat float64, z float64) (tile *geois.Tile) {
	scaleZ := math.Exp2(z)
	tileTempX := (lon + 180.0) / 360.0 * scaleZ
	tileTempY := math.Log(math.Tan(lat*math.Pi/180*0.5+0.25*math.Pi)) / (2.0 * math.Pi)
	//
	tileX := math.Floor(tileTempX)
	tileY := math.Floor((0.5 - tileTempY) * scaleZ)

	return &geois.Tile{
		X: int(tileX),
		Y: int(tileY),
		Z: int(z),
	}
}

// LonLat2TilePoint 经纬度转成瓦片中的一个点
func LonLat2TilePoint(lon float64, lat float64, z float64) (tilePoint *geois.TilePoint) {
	scaleZ := math.Exp2(z)
	tileTempX := (lon + 180.0) / 360.0 * scaleZ
	tileTempY := math.Log(math.Tan(lat*math.Pi/180*0.5+0.25*math.Pi)) / (2.0 * math.Pi)
	//
	tileX := math.Floor(tileTempX)
	tileY := math.Floor((0.5 - tileTempY) * scaleZ)
	//
	pixelX := int(tileTempX*256.0) % 256
	pixelY := int((1.0-tileTempY)*scaleZ*256.0) % 256
	//
	offsetPoint := &grap.Point{
		X: float64(pixelX),
		Y: float64(pixelY),
	}

	tilePoint = &geois.TilePoint{
		Tile: geois.Tile{
			X: int(tileX),
			Y: int(tileY),
			Z: int(z),
		},
		Offset: offsetPoint,
	}

	return
}

// Tile2LonLat 瓦片转经纬度(左上角)
func Tile2LonLat(tile *geois.Tile) *grap.Point {
	x, y, z := tile.X, tile.Y, tile.Z
	n := math.Pi - 2*math.Pi*float64(y)/math.Pow(float64(2), float64(z))

	return &grap.Point{
		X: float64(x)/math.Exp2(float64(z))*360 - 180,
		Y: (R2D * math.Atan(0.5*(math.Exp(n)-math.Exp(-1.0*n)))),
	}
}
