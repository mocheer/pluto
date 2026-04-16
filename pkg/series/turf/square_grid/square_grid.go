package square_grid

import (
	"errors"
	"math"

	"github.com/mocheer/xena/pkg/gm"
)

type SquareGridOptions struct {
	CellSize float64
	BBox     gm.BBox         // 用于范围生成，这个范围小于最终的范围
	Mask     gm.MultiPolygon // 用于裁剪
}

type SquareGrid struct {
	Data     []float64
	Columns  int
	Rows     int
	CellSize float64
	Xlim     [2]float64
	Ylim     [2]float64
	BBox     gm.BBox
}

func New() *SquareGrid {
	return &SquareGrid{}
}

// NewWithOptions
func NewWithOptions(options SquareGridOptions) (*SquareGrid, error) {
	// 参数检查
	cellSide := options.CellSize
	if cellSide <= 0 {
		return nil, errors.New("cellSide must be positive")
	}
	bbox := options.BBox
	width := bbox.Width()
	height := bbox.Height()

	columns := int(math.Ceil(width / options.CellSize))
	rows := int(math.Ceil(height / options.CellSize))

	// 调整网格起点
	deltaX := (width - float64(columns)*options.CellSize) / 2
	deltaY := (height - float64(rows)*options.CellSize) / 2
	// 边界框宽度和高度
	data := make([]float64, 0, columns*rows)

	grid := &SquareGrid{
		Data:     data,
		Columns:  columns,
		Rows:     rows,
		Xlim:     [2]float64{bbox.MinX() - deltaX, bbox.MaxX() + deltaX},
		Ylim:     [2]float64{bbox.MinY() - deltaY, bbox.MaxY() + deltaY},
		CellSize: options.CellSize,
	}
	return grid, nil
}
