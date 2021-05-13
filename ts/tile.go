package ts

// Tile 地图瓦片
type Tile struct {
	X, Y, Z int
}

type TilePoint struct {
	Tile
	Offset *Point
}
