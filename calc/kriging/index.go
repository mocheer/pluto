package kriging

import (
	"github.com/liuvigongzuoshi/go-kriging/ordinarykriging"
)

type k struct {
	Variogram *ordinarykriging.Variogram
}

func New(values, x, y []float64) *k {
	return &k{
		Variogram: ordinarykriging.NewOrdinary(values, x, y),
	}
}

// Train
func (m *k) Train(modelType string) error {
	_, err := m.Variogram.Train(ordinarykriging.ModelType(modelType), 0, 100)
	return err
}

// Grid 生成网格数据,这个网格数据是从上到下，再从左到右的
// cellSize 一般是0.01(地图应用)
func (m *k) Grid(polygon [][2]float64, cellSize float64) *ordinarykriging.GridMatrices {
	ring := make(ordinarykriging.Ring, len(polygon))
	for i, p := range polygon {
		ring[i] = ordinarykriging.Point{p[0], p[1]}
	}
	coordinates := ordinarykriging.PolygonCoordinates{ring}
	return m.Variogram.Grid(coordinates, cellSize)
}

// Contour
func (m *k) Contour(xWidth int, yWidth int) *ordinarykriging.ContourRectangle {
	return m.Variogram.Contour(xWidth, yWidth)
}

// SavePng 将网格数据保存为图片
func (m *k) SaveGridPng(grid *ordinarykriging.GridMatrices, dst string, xWidth int, yWidth int, levelColors []ordinarykriging.GridLevelColor) error {
	ctx := m.Variogram.Plot(grid, xWidth, yWidth, grid.Xlim, grid.Ylim, levelColors)
	return ctx.SavePNG(dst)
}
