package ds_ewkb

// Package ewkb 实现了扩展的 WKB (Well Known Binary) 编码和解码。
// 参见 https://github.com/postgis/postgis/blob/2.1.0/doc/ZMSgeoms.txt。
//
// 如果您要编码 EWKB 几何体以发送到 PostgreSQL/PostGIS，
// 则必须在传递给 sql.Open 的数据源名称中指定 binary_parameters=yes。
import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
)

// ==================== 类型定义 ====================

// Layout 定义了几何体的坐标布局（维度）
type Layout int

const (
	XY   Layout = iota // XY: 2D 几何体 (x, y)
	XYZ                // XYZ: 3D 几何体 (x, y, z)
	XYM                // XYM: 带测量值的 2D 几何体 (x, y, m)
	XYZM               // XYZM: 带测量值的 3D 几何体 (x, y, z, m)
)

// Stride 返回该布局的坐标维度数量
func (l Layout) Stride() int {
	switch l {
	case XY:
		return 2
	case XYZ:
		return 3
	case XYM:
		return 3
	case XYZM:
		return 4
	default:
		return 0
	}
}

// GeometryType 表示几何体类型
type GeometryType int

const (
	PointType              GeometryType = 1
	LineStringType         GeometryType = 2
	PolygonType            GeometryType = 3
	MultiPointType         GeometryType = 4
	MultiLineStringType    GeometryType = 5
	MultiPolygonType       GeometryType = 6
	GeometryCollectionType GeometryType = 7
)

// Geometry 表示一个几何体
type Geometry struct {
	Type   GeometryType // 几何体类型
	Layout Layout       // 坐标布局
	SRID   int          // 空间参考系统ID

	// 坐标数据，根据几何体类型不同而不同：
	// - Point: []float64 (长度由布局决定)
	// - LineString: [][]float64 (每个元素是点)
	// - Polygon: [][][]float64 (第一维是环，第二维是点)
	// - MultiPoint: [][]float64 (每个元素是点)
	// - MultiLineString: [][][]float64 (第一维是线串，第二维是点)
	// - MultiPolygon: [][][][]float64 (第一维是多边形，第二维是环，第三维是点)
	// - GeometryCollection: []Geometry
	Coords interface{}
}

// ==================== 常量定义 ====================

const (
	ewkbZ    uint32 = 0x80000000 // Z坐标标志位
	ewkbM    uint32 = 0x40000000 // M坐标（测量值）标志位
	ewkbSRID uint32 = 0x20000000 // SRID标志位
)

// 字节顺序标识
const (
	XDRID byte = 0 // 大端序
	NDRID byte = 1 // 小端序
)

// ==================== 错误定义 ====================

var (
	// ErrUnknownByteOrder 表示未知的字节顺序
	ErrUnknownByteOrder = errors.New("ewkb: 未知字节顺序")
	// ErrUnknownType 表示未知的几何体类型
	ErrUnknownType = errors.New("ewkb: 未知几何体类型")
	// ErrUnsupportedType 表示不支持的几何体类型
	ErrUnsupportedType = errors.New("ewkb: 不支持的几何体类型")
	// ErrUnexpectedType 表示意外的几何体类型
	ErrUnexpectedType = errors.New("ewkb: 意外的几何体类型")
	// ErrGeometryTooLarge 表示几何体太大
	ErrGeometryTooLarge = errors.New("ewkb: 几何体太大")
	// ErrUnsupportedByteOrder 表示不支持的字节顺序
	ErrUnsupportedByteOrder = errors.New("ewkb: 不支持的字节顺序")
	// ErrUnsupportedLayout 表示不支持的布局
	ErrUnsupportedLayout = errors.New("ewkb: 不支持的布局")
)

// ==================== 辅助函数 ====================

// readByte 从 Reader 读取一个字节
func readByte(r io.Reader) (byte, error) {
	var b [1]byte
	_, err := r.Read(b[:])
	return b[0], err
}

// readUInt32 从 Reader 以指定字节顺序读取一个 uint32
func readUInt32(r io.Reader, byteOrder binary.ByteOrder) (uint32, error) {
	var val uint32
	err := binary.Read(r, byteOrder, &val)
	return val, err
}

// readFloat64 从 Reader 以指定字节顺序读取一个 float64
func readFloat64(r io.Reader, byteOrder binary.ByteOrder) (float64, error) {
	var val float64
	err := binary.Read(r, byteOrder, &val)
	return val, err
}

// writeFloat64 将 float64 以指定字节顺序写入 Writer
func writeFloat64(w io.Writer, byteOrder binary.ByteOrder, val float64) error {
	return binary.Write(w, byteOrder, val)
}

// readFlatCoords0 读取点坐标（0维数组）
func readFlatCoords0(r io.Reader, byteOrder binary.ByteOrder, stride int) ([]float64, error) {
	if stride <= 0 {
		return nil, fmt.Errorf("无效的 stride: %d", stride)
	}

	coords := make([]float64, stride)
	for i := 0; i < stride; i++ {
		coord, err := readFloat64(r, byteOrder)
		if err != nil {
			return nil, err
		}
		coords[i] = coord
	}
	return coords, nil
}

// readFlatCoords1 读取线串坐标（1维数组）
func readFlatCoords1(r io.Reader, byteOrder binary.ByteOrder, stride int) ([][]float64, error) {
	numPoints, err := readUInt32(r, byteOrder)
	if err != nil {
		return nil, err
	}

	if numPoints == 0 {
		return [][]float64{}, nil
	}

	points := make([][]float64, numPoints)
	for i := 0; i < int(numPoints); i++ {
		point, err := readFlatCoords0(r, byteOrder, stride)
		if err != nil {
			return nil, err
		}
		points[i] = point
	}
	return points, nil
}

// readFlatCoords2 读取多边形坐标（2维数组）
func readFlatCoords2(r io.Reader, byteOrder binary.ByteOrder, stride int) ([][][]float64, error) {
	numRings, err := readUInt32(r, byteOrder)
	if err != nil {
		return nil, err
	}

	if numRings == 0 {
		return [][][]float64{}, nil
	}

	rings := make([][][]float64, numRings)
	for i := 0; i < int(numRings); i++ {
		ring, err := readFlatCoords1(r, byteOrder, stride)
		if err != nil {
			return nil, err
		}
		rings[i] = ring
	}
	return rings, nil
}

// writeFlatCoords0 写入点坐标
func writeFlatCoords0(w io.Writer, byteOrder binary.ByteOrder, coords []float64) error {
	for _, coord := range coords {
		if err := writeFloat64(w, byteOrder, coord); err != nil {
			return err
		}
	}
	return nil
}

// writeFlatCoords1 写入线串坐标
func writeFlatCoords1(w io.Writer, byteOrder binary.ByteOrder, points [][]float64, stride int) error {
	// 写入点数
	numPoints := uint32(len(points))
	if err := binary.Write(w, byteOrder, numPoints); err != nil {
		return err
	}

	// 写入每个点
	for _, point := range points {
		if len(point) != stride {
			return fmt.Errorf("点坐标维度 %d 与 stride %d 不匹配", len(point), stride)
		}
		if err := writeFlatCoords0(w, byteOrder, point); err != nil {
			return err
		}
	}
	return nil
}

// writeFlatCoords2 写入多边形坐标
func writeFlatCoords2(w io.Writer, byteOrder binary.ByteOrder, rings [][][]float64, ends []int, stride int) error {
	// 写入环数
	numRings := uint32(len(rings))
	if err := binary.Write(w, byteOrder, numRings); err != nil {
		return err
	}

	// 写入每个环
	for _, ring := range rings {
		if err := writeFlatCoords1(w, byteOrder, ring, stride); err != nil {
			return err
		}
	}
	return nil
}

// writeEmptyPointAsNaN 将空点写为 NaN 值
func writeEmptyPointAsNaN(w io.Writer, byteOrder binary.ByteOrder, stride int) error {
	nan := math.NaN()
	for i := 0; i < stride; i++ {
		if err := writeFloat64(w, byteOrder, nan); err != nil {
			return err
		}
	}
	return nil
}

// ==================== 公共函数 ====================

// Read 从 Reader 读取任意几何体
func Read(r io.Reader) (*Geometry, error) {
	// 读取字节顺序
	ewkbByteOrder, err := readByte(r)
	if err != nil {
		return nil, err
	}

	var byteOrder binary.ByteOrder
	switch ewkbByteOrder {
	case XDRID:
		byteOrder = binary.BigEndian
	case NDRID:
		byteOrder = binary.LittleEndian
	default:
		return nil, fmt.Errorf("%w: %d", ErrUnknownByteOrder, ewkbByteOrder)
	}

	// 读取几何体类型
	ewkbGeometryType, err := readUInt32(r, byteOrder)
	if err != nil {
		return nil, err
	}

	// 确定布局
	var layout Layout
	switch ewkbGeometryType & (ewkbZ | ewkbM) {
	case 0:
		layout = XY
	case ewkbZ:
		layout = XYZ
	case ewkbM:
		layout = XYM
	case ewkbZ | ewkbM:
		layout = XYZM
	default:
		return nil, fmt.Errorf("%w: %d", ErrUnknownType, ewkbGeometryType&(ewkbZ|ewkbM))
	}

	// 读取 SRID（如果存在）
	var srid uint32
	if ewkbGeometryType&ewkbSRID != 0 {
		srid, err = readUInt32(r, byteOrder)
		if err != nil {
			return nil, err
		}
	}

	// 根据几何体类型读取几何体数据
	geometryType := ewkbGeometryType &^ (ewkbZ | ewkbM | ewkbSRID)

	switch geometryType {
	case uint32(PointType):
		flatCoords, err := readFlatCoords0(r, byteOrder, layout.Stride())
		if err != nil {
			return nil, err
		}
		// 检查是否是空点（所有坐标都是 NaN）
		isEmpty := true
		for _, coord := range flatCoords {
			if !math.IsNaN(coord) {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			flatCoords = []float64{}
		}
		return &Geometry{
			Type:   PointType,
			Layout: layout,
			SRID:   int(srid),
			Coords: flatCoords,
		}, nil

	case uint32(LineStringType):
		flatCoords, err := readFlatCoords1(r, byteOrder, layout.Stride())
		if err != nil {
			return nil, err
		}
		return &Geometry{
			Type:   LineStringType,
			Layout: layout,
			SRID:   int(srid),
			Coords: flatCoords,
		}, nil

	case uint32(PolygonType):
		flatCoords, err := readFlatCoords2(r, byteOrder, layout.Stride())
		if err != nil {
			return nil, err
		}
		// 计算每个环的结束索引
		ends := make([]int, len(flatCoords))
		for i, ring := range flatCoords {
			ends[i] = len(ring)
		}
		return &Geometry{
			Type:   PolygonType,
			Layout: layout,
			SRID:   int(srid),
			Coords: flatCoords,
		}, nil

	case uint32(MultiPointType):
		n, err := readUInt32(r, byteOrder)
		if err != nil {
			return nil, err
		}

		points := make([][]float64, 0, n)
		for i := 0; i < int(n); i++ {
			g, err := Read(r)
			if err != nil {
				return nil, err
			}
			if g.Type != PointType {
				return nil, fmt.Errorf("%w: 期望点类型，实际为 %d", ErrUnexpectedType, g.Type)
			}
			pointCoords, ok := g.Coords.([]float64)
			if !ok {
				return nil, fmt.Errorf("多点中的点坐标类型错误")
			}
			points = append(points, pointCoords)
		}

		return &Geometry{
			Type:   MultiPointType,
			Layout: layout,
			SRID:   int(srid),
			Coords: points,
		}, nil

	case uint32(MultiLineStringType):
		n, err := readUInt32(r, byteOrder)
		if err != nil {
			return nil, err
		}

		lineStrings := make([][][]float64, 0, n)
		for i := 0; i < int(n); i++ {
			g, err := Read(r)
			if err != nil {
				return nil, err
			}
			if g.Type != LineStringType {
				return nil, fmt.Errorf("%w: 期望线串类型，实际为 %d", ErrUnexpectedType, g.Type)
			}
			lineCoords, ok := g.Coords.([][]float64)
			if !ok {
				return nil, fmt.Errorf("多线串中的线串坐标类型错误")
			}
			lineStrings = append(lineStrings, lineCoords)
		}

		return &Geometry{
			Type:   MultiLineStringType,
			Layout: layout,
			SRID:   int(srid),
			Coords: lineStrings,
		}, nil

	case uint32(MultiPolygonType):
		n, err := readUInt32(r, byteOrder)
		if err != nil {
			return nil, err
		}

		polygons := make([][][][]float64, 0, n)
		for i := 0; i < int(n); i++ {
			g, err := Read(r)
			if err != nil {
				return nil, err
			}
			if g.Type != PolygonType {
				return nil, fmt.Errorf("%w: 期望多边形类型，实际为 %d", ErrUnexpectedType, g.Type)
			}
			polyCoords, ok := g.Coords.([][][]float64)
			if !ok {
				return nil, fmt.Errorf("多多边形中的多边形坐标类型错误")
			}
			polygons = append(polygons, polyCoords)
		}

		return &Geometry{
			Type:   MultiPolygonType,
			Layout: layout,
			SRID:   int(srid),
			Coords: polygons,
		}, nil

	case uint32(GeometryCollectionType):
		n, err := readUInt32(r, byteOrder)
		if err != nil {
			return nil, err
		}

		geometries := make([]Geometry, 0, n)
		for i := 0; i < int(n); i++ {
			g, err := Read(r)
			if err != nil {
				return nil, err
			}
			geometries = append(geometries, *g)
		}

		return &Geometry{
			Type:   GeometryCollectionType,
			Layout: layout,
			SRID:   int(srid),
			Coords: geometries,
		}, nil

	default:
		return nil, fmt.Errorf("%w: %d", ErrUnsupportedType, geometryType)
	}
}

// Unmarshal 从字节切片解析任意几何体
func Unmarshal(data []byte) (*Geometry, error) {
	return Read(bytes.NewBuffer(data))
}

// Write 将几何体写入 Writer
func Write(w io.Writer, byteOrder binary.ByteOrder, g *Geometry) error {
	// 写入字节顺序标识
	var ewkbByteOrder byte
	switch byteOrder {
	case binary.BigEndian:
		ewkbByteOrder = XDRID
	case binary.LittleEndian:
		ewkbByteOrder = NDRID
	default:
		return ErrUnsupportedByteOrder
	}

	if err := binary.Write(w, byteOrder, ewkbByteOrder); err != nil {
		return err
	}

	// 确定几何体类型和标志位
	var ewkbGeometryType uint32

	switch g.Type {
	case PointType:
		ewkbGeometryType = uint32(PointType)
	case LineStringType:
		ewkbGeometryType = uint32(LineStringType)
	case PolygonType:
		ewkbGeometryType = uint32(PolygonType)
	case MultiPointType:
		ewkbGeometryType = uint32(MultiPointType)
	case MultiLineStringType:
		ewkbGeometryType = uint32(MultiLineStringType)
	case MultiPolygonType:
		ewkbGeometryType = uint32(MultiPolygonType)
	case GeometryCollectionType:
		ewkbGeometryType = uint32(GeometryCollectionType)
	default:
		return fmt.Errorf("%w: %v", ErrUnsupportedType, g.Type)
	}

	// 添加布局标志位
	switch g.Layout {
	case XY:
		// 无额外标志
	case XYZ:
		ewkbGeometryType |= ewkbZ
	case XYM:
		ewkbGeometryType |= ewkbM
	case XYZM:
		ewkbGeometryType |= ewkbZ | ewkbM
	default:
		return fmt.Errorf("%w: %v", ErrUnsupportedLayout, g.Layout)
	}

	// 添加 SRID 标志位（如果需要）
	if g.SRID != 0 {
		ewkbGeometryType |= ewkbSRID
	}

	// 写入几何体类型
	if err := binary.Write(w, byteOrder, ewkbGeometryType); err != nil {
		return err
	}

	// 写入 SRID（如果需要）
	if ewkbGeometryType&ewkbSRID != 0 {
		if err := binary.Write(w, byteOrder, uint32(g.SRID)); err != nil {
			return err
		}
	}

	// 根据几何体类型写入坐标数据
	stride := g.Layout.Stride()

	switch g.Type {
	case PointType:
		coords, ok := g.Coords.([]float64)
		if !ok {
			return fmt.Errorf("点坐标类型错误")
		}

		if len(coords) == 0 {
			// 空点
			return writeEmptyPointAsNaN(w, byteOrder, stride)
		}
		return writeFlatCoords0(w, byteOrder, coords)

	case LineStringType:
		coords, ok := g.Coords.([][]float64)
		if !ok {
			return fmt.Errorf("线串坐标类型错误")
		}
		return writeFlatCoords1(w, byteOrder, coords, stride)

	case PolygonType:
		coords, ok := g.Coords.([][][]float64)
		if !ok {
			return fmt.Errorf("多边形坐标类型错误")
		}
		// 为 writeFlatCoords2 创建 ends 参数
		ends := make([]int, len(coords))
		for i, ring := range coords {
			ends[i] = len(ring)
		}
		return writeFlatCoords2(w, byteOrder, coords, ends, stride)

	case MultiPointType:
		points, ok := g.Coords.([][]float64)
		if !ok {
			return fmt.Errorf("多点坐标类型错误")
		}

		// 写入点数
		n := uint32(len(points))
		if err := binary.Write(w, byteOrder, n); err != nil {
			return err
		}

		// 写入每个点
		for _, point := range points {
			pointGeom := &Geometry{
				Type:   PointType,
				Layout: g.Layout,
				SRID:   g.SRID,
				Coords: point,
			}
			if err := Write(w, byteOrder, pointGeom); err != nil {
				return err
			}
		}
		return nil

	case MultiLineStringType:
		lineStrings, ok := g.Coords.([][][]float64)
		if !ok {
			return fmt.Errorf("多线串坐标类型错误")
		}

		// 写入线串数
		n := uint32(len(lineStrings))
		if err := binary.Write(w, byteOrder, n); err != nil {
			return err
		}

		// 写入每个线串
		for _, lineString := range lineStrings {
			lineGeom := &Geometry{
				Type:   LineStringType,
				Layout: g.Layout,
				SRID:   g.SRID,
				Coords: lineString,
			}
			if err := Write(w, byteOrder, lineGeom); err != nil {
				return err
			}
		}
		return nil

	case MultiPolygonType:
		polygons, ok := g.Coords.([][][][]float64)
		if !ok {
			return fmt.Errorf("多多边形坐标类型错误")
		}

		// 写入多边形数
		n := uint32(len(polygons))
		if err := binary.Write(w, byteOrder, n); err != nil {
			return err
		}

		// 写入每个多边形
		for _, polygon := range polygons {
			polyGeom := &Geometry{
				Type:   PolygonType,
				Layout: g.Layout,
				SRID:   g.SRID,
				Coords: polygon,
			}
			if err := Write(w, byteOrder, polyGeom); err != nil {
				return err
			}
		}
		return nil

	case GeometryCollectionType:
		geometries, ok := g.Coords.([]Geometry)
		if !ok {
			return fmt.Errorf("几何体集合坐标类型错误")
		}

		// 写入几何体数
		n := uint32(len(geometries))
		if err := binary.Write(w, byteOrder, n); err != nil {
			return err
		}

		// 写入每个几何体
		for _, geom := range geometries {
			if err := Write(w, byteOrder, &geom); err != nil {
				return err
			}
		}
		return nil

	default:
		return fmt.Errorf("%w: %v", ErrUnsupportedType, g.Type)
	}
}

// Marshal 将几何体编码为字节切片
func Marshal(g *Geometry, byteOrder binary.ByteOrder) ([]byte, error) {
	w := bytes.NewBuffer(nil)
	if err := Write(w, byteOrder, g); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

// ==================== 辅助函数 ====================

// NewPoint 创建点几何体
func NewPoint(layout Layout, coords []float64, srid int) *Geometry {
	return &Geometry{
		Type:   PointType,
		Layout: layout,
		SRID:   srid,
		Coords: coords,
	}
}

// NewLineString 创建线串几何体
func NewLineString(layout Layout, coords [][]float64, srid int) *Geometry {
	return &Geometry{
		Type:   LineStringType,
		Layout: layout,
		SRID:   srid,
		Coords: coords,
	}
}

// NewPolygon 创建多边形几何体
func NewPolygon(layout Layout, coords [][][]float64, srid int) *Geometry {
	return &Geometry{
		Type:   PolygonType,
		Layout: layout,
		SRID:   srid,
		Coords: coords,
	}
}

// NewMultiPoint 创建多点几何体
func NewMultiPoint(layout Layout, coords [][]float64, srid int) *Geometry {
	return &Geometry{
		Type:   MultiPointType,
		Layout: layout,
		SRID:   srid,
		Coords: coords,
	}
}

// NewMultiLineString 创建多线串几何体
func NewMultiLineString(layout Layout, coords [][][]float64, srid int) *Geometry {
	return &Geometry{
		Type:   MultiLineStringType,
		Layout: layout,
		SRID:   srid,
		Coords: coords,
	}
}

// NewMultiPolygon 创建多多边形几何体
func NewMultiPolygon(layout Layout, coords [][][][]float64, srid int) *Geometry {
	return &Geometry{
		Type:   MultiPolygonType,
		Layout: layout,
		SRID:   srid,
		Coords: coords,
	}
}

// NewGeometryCollection 创建几何体集合
func NewGeometryCollection(layout Layout, geometries []Geometry, srid int) *Geometry {
	return &Geometry{
		Type:   GeometryCollectionType,
		Layout: layout,
		SRID:   srid,
		Coords: geometries,
	}
}
