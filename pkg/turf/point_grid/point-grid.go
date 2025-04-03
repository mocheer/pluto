package point_grid

import (
	"errors"
	"math"

	"github.com/mocheer/pluto/pkg/turf/distance"
	"github.com/mocheer/xena/pkg/gm"
)

// TODO
// turf支持mask，目前打算将mask放到前端遮罩实现
// turf返回结果是geojson，这里直接用更原始的数据结构
// turf支持传递properties，这里不实现

// 新增了 columns 和rows输出，方便栅格统计

type PointGridOptions struct {
	Units    string
	CellSize float64
	BBox     gm.BBox
	Mask     gm.MultiPolygon // 用于裁剪
}

type PointGrid struct {
	Points     []gm.Cartographic
	Columns    int
	Rows       int
	CellWidth  float64
	CellHeight float64
	Xlim       [2]float64
	Ylim       [2]float64
}

func New() *PointGrid {
	return &PointGrid{}
}

// PointGrid 创建一个点网格
func NewWithOptions(options PointGridOptions) (*PointGrid, error) {
	// 参数检查
	cellSide := options.CellSize
	if cellSide <= 0 {
		return nil, errors.New("cellSide must be positive")
	}
	// 边界框范围
	bbox := options.BBox
	west, south, east, north := bbox[0], bbox[1], bbox[2], bbox[3]
	// 边界框宽度和高度
	bboxWidth := bbox.Width()
	bboxHeight := bbox.Height()

	// 计算网格行列数
	columns := int(math.Ceil(bboxWidth / cellSide))
	rows := int(math.Ceil(bboxHeight / cellSide))

	// 调整网格起点
	deltaX := (bboxWidth - float64(columns-1)*cellSide) / 2
	deltaY := (bboxHeight - float64(rows-1)*cellSide) / 2

	xlim := [2]float64{west + deltaX, east - deltaX}
	ylim := [2]float64{south + deltaY, north - deltaY}

	// 创建网格点
	results := make([]gm.Cartographic, 0, columns*rows)
	currentX := xlim[0]
	for currentX <= east {
		currentY := ylim[0]
		for currentY <= north {
			cellPt := gm.LonLat{currentX, currentY}.ToCartographic()
			results = append(results, cellPt)
			currentY += cellSide
		}
		currentX += cellSide
	}
	grid := &PointGrid{
		Points:     results,
		Columns:    columns,
		Rows:       rows,
		Xlim:       xlim,
		Ylim:       ylim,
		CellWidth:  cellSide,
		CellHeight: cellSide,
	}
	return grid, nil
}

// PointGrid 创建一个点网格
func NewWithOptionsWithCellFraction(options PointGridOptions) (*PointGrid, error) {
	// 参数检查
	cellSide := options.CellSize
	if cellSide <= 0 {
		return nil, errors.New("cellSide must be positive")
	}
	// 边界框范围
	bbox := options.BBox
	west, south, east, north := bbox[0], bbox[1], bbox[2], bbox[3]
	// 边界框宽度和高度
	bboxWidth := bbox.Width()
	bboxHeight := bbox.Height()
	westSouth := bbox.WestSouth()

	// 计算单元格宽度和高度
	xFraction := cellSide / distance.DistanceLonLat(westSouth, gm.LonLat{east, south}, options.Units)
	cellWidth := xFraction * bboxWidth
	yFraction := cellSide / distance.DistanceLonLat(westSouth, gm.LonLat{west, north}, options.Units)
	cellHeight := yFraction * bboxHeight

	// 计算网格行列数
	columns := int(math.Ceil(bboxWidth / cellWidth))
	rows := int(math.Ceil(bboxHeight / cellHeight))

	// 调整网格起点
	deltaX := (bboxWidth - float64(columns-1)*cellWidth) / 2
	deltaY := (bboxHeight - float64(rows-1)*cellHeight) / 2

	xlim := [2]float64{west + deltaX, east - deltaX}
	ylim := [2]float64{south + deltaY, north - deltaY}

	// 创建网格点
	results := make([]gm.Cartographic, 0, columns*rows)
	currentX := xlim[0]
	for currentX <= east {
		currentY := ylim[0]
		for currentY <= north {
			cellPt := gm.LonLat{currentX, currentY}.ToCartographic()
			results = append(results, cellPt)
			currentY += cellHeight
		}
		currentX += cellWidth
	}
	grid := &PointGrid{
		Points:     results,
		Columns:    columns,
		Rows:       rows,
		Xlim:       xlim,
		Ylim:       ylim,
		CellWidth:  cellWidth,
		CellHeight: cellHeight,
	}
	return grid, nil
}
