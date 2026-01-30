package ds_raster

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/samber/lo"
)

// https://github.com/nathancahill/wkb/blob/master/wkb.go
// https://github.com/paulmach/orb/tree/master/encoding/ewkb
// https://github.com/fschutt/wkb-raster/blob/master/src/lib.rs#L1146
// https://github.com/nathancahill/wkb-raster/blob/master/wkb_raster.py
// https://github.com/ihmeuw/wkb-raster/blob/master/src/constants.ts

// 像素类型映射
// []string{"1BB","2BUI","4BUI","8SI","8BUI","16BSI","16BUI","32BSI","32BUI","32BF","64BF"}
// 像素类型映射有点问题的是，没有9，会跳过9直接到10和11 => 有点奇怪
// 计算机内存最小寻址单位是字节，所以最低是1字节
type PixelType interface {
	uint8 | int8 | uint16 | int16 | uint32 | int32 | float32 | float64
}

// Header WKB栅格头信息结构体
type Header struct {
	Version uint16  // 格式版本（0表示当前结构）
	Bands   uint16  // 波段数量
	ScaleX  float64 // 像素宽度（地理单位）
	ScaleY  float64 // 像素高度（地理单位）
	IpX     float64 // 左上角像素左上角的X坐标（地理单位）
	IpY     float64 // 左上角像素左上角的Y坐标（地理单位）
	SkewX   float64 // Y轴旋转
	SkewY   float64 // X轴旋转
	Srid    int32   // 空间参考ID
	Width   uint16  // 像素列数，最高65536
	Height  uint16  // 像素行数，最高65536
}

// BandHeader 波段头信息（1字节的位字段）
type BandHeader struct {
	IsOffline      bool    // 第1位：是否为离线数据
	HasNoDataValue bool    // 第2位：是否有无数据值
	IsNoDataValue  bool    // 第3位：是否所有值都是无数据值（脏标志）
	Reserved       bool    // 第4位：保留位（当前版本未使用）
	PixelType      int     // 第5-8位：像素类型索引（0-10）
	NoData         float64 //无数据值
}

// Band 波段完整信息
type Band struct {
	Header BandHeader // 波段头信息
	Data   []float64  // 波段数据
	Path   string     // 离线数据文件路径（仅当IsOffline为true时有效）
}

// WKBRaster 完整的WKB栅格数据结构
type WKBRaster struct {
	Header Header // 栅格头信息
	Bands  []Band // 波段列表
}

// readAsFloat64Slice
func readAsFloat64Slice[T PixelType](v []T) []float64 {
	return lo.Map(v, func(item T, _ int) float64 {
		return float64(item)
	})
}

// ReadWKBRaster 从WKB格式读取栅格数据
func ReadWKBRaster(r io.Reader) (*WKBRaster, error) {
	// 读取字节序标识
	var endianByte byte
	if err := binary.Read(r, binary.LittleEndian, &endianByte); err != nil {
		return nil, fmt.Errorf("读取字节序失败: %v", err)
	}

	// 确定字节序
	var byteOrder binary.ByteOrder
	switch endianByte {
	case 0:
		byteOrder = binary.BigEndian // XDR/大端序
	case 1:
		byteOrder = binary.LittleEndian // NDR/小端序
	default:
		return nil, fmt.Errorf("无效的字节序标识: %d", endianByte)
	}

	// 读取头信息
	var header Header
	if err := binary.Read(r, byteOrder, &header); err != nil {
		return nil, fmt.Errorf("读取栅格头信息失败: %v", err)
	}

	// 验证头信息
	if err := validateHeader(&header); err != nil {
		return nil, fmt.Errorf("栅格头信息验证失败: %v", err)
	}

	// 创建WKBRaster结构
	raster := &WKBRaster{
		Header: header,
		Bands:  make([]Band, 0, header.Bands),
	}

	// 读取每个波段
	for i := uint16(0); i < header.Bands; i++ {
		band, err := readBand(r, byteOrder, int(header.Width), int(header.Height))
		if err != nil {
			return nil, fmt.Errorf("读取波段%d失败: %v", i, err)
		}
		raster.Bands = append(raster.Bands, *band)
	}

	return raster, nil
}

// validateHeader 验证栅格头信息的有效性
func validateHeader(header *Header) error {
	if header.Version != 0 {
		return fmt.Errorf("不支持的WKB版本: %d", header.Version)
	}
	if header.Bands == 0 {
		return fmt.Errorf("波段数量不能为0")
	}
	if header.Width == 0 || header.Height == 0 {
		return fmt.Errorf("栅格尺寸无效: %dx%d", header.Width, header.Height)
	}
	if header.ScaleX == 0 || header.ScaleY == 0 {
		return fmt.Errorf("像素大小无效: scaleX=%f, scaleY=%f", header.ScaleX, header.ScaleY)
	}
	return nil
}

// parseBandHeader 解析波段头字节
func parseBandHeader(headerByte byte) BandHeader {
	return BandHeader{
		IsOffline:      (headerByte & 0x80) != 0, // 第1位
		HasNoDataValue: (headerByte & 0x40) != 0, // 第2位
		IsNoDataValue:  (headerByte & 0x20) != 0, // 第3位
		Reserved:       (headerByte & 0x10) != 0, // 第4位
		PixelType:      int(headerByte & 0x0F),   // 第5-8位，需要注意的是这里面没有9
	}
}

// buildBandHeader 构建波段头字节
func buildBandHeader(header BandHeader) byte {
	var headerByte byte

	if header.IsOffline {
		headerByte |= 0x80 // 设置第1位
	}
	if header.HasNoDataValue {
		headerByte |= 0x40 // 设置第2位
	}
	if header.IsNoDataValue {
		headerByte |= 0x20 // 设置第3位
	}
	if header.Reserved {
		headerByte |= 0x10 // 设置第4位
	}

	// 设置像素类型（限制在0-15范围内）
	headerByte |= byte(header.PixelType & 0x0F)

	return headerByte
}

// readBand 读取单个波段数据
func readBand(r io.Reader, byteOrder binary.ByteOrder, width, height int) (*Band, error) {
	// 读取波段头字节
	var headerByte byte
	if err := binary.Read(r, byteOrder, &headerByte); err != nil {
		return nil, fmt.Errorf("读取波段头字节失败: %v", err)
	}

	// 解析波段头
	bandHeader := parseBandHeader(headerByte)

	// 检查像素类型是否有效
	if bandHeader.PixelType < 0 || bandHeader.PixelType > 11 {
		return nil, fmt.Errorf("无效的像素类型: %d", bandHeader.PixelType)
	}
	// 读取无数据值
	if bandHeader.HasNoDataValue {

		noDataValue, err := readValueByPixelType(r, byteOrder, bandHeader.PixelType)
		if err != nil {
			return nil, fmt.Errorf("读取无数据值失败: %v", err)
		}
		bandHeader.NoData = noDataValue
	}

	// 创建Band结构
	band := &Band{
		Header: bandHeader,
	}

	if bandHeader.IsOffline {
		// 读取离线数据
		if err := readOfflineBand(r, byteOrder, band); err != nil {
			return nil, err
		}
	} else {
		// 读取像素数据
		if err := readPixelData(r, byteOrder, bandHeader.PixelType, width, height, band); err != nil {
			return nil, err
		}
	}

	return band, nil
}

// readOfflineBand 读取离线波段数据
func readOfflineBand(r io.Reader, byteOrder binary.ByteOrder, band *Band) error {
	// 读取波段编号
	var bandNum byte
	if err := binary.Read(r, byteOrder, &bandNum); err != nil {
		return fmt.Errorf("读取波段编号失败: %v", err)
	}

	// 读取以null结尾的文件路径
	var pathBuilder strings.Builder
	for {
		var ch byte
		if err := binary.Read(r, byteOrder, &ch); err != nil {
			return fmt.Errorf("读取文件路径失败: %v", err)
		}
		if ch == 0 {
			break
		}
		pathBuilder.WriteByte(ch)
	}

	band.Path = pathBuilder.String()
	return nil
}

// readPixelData 读取像素数据
func readPixelData(r io.Reader, byteOrder binary.ByteOrder, pixelType int,
	width, height int, band *Band) error {
	totalPixels := width * height
	data, err := readValuesByPixelType(r, byteOrder, pixelType, totalPixels)
	if err != nil {
		return err
	}
	band.Data = data
	// for i := 0; i < totalPixels; i++ {
	// 	val, err := readValueByPixelType(r, byteOrder,pixelType)
	// 	if err != nil {
	// 		return fmt.Errorf("读取像素%d失败: %v", i, err)
	// 	}
	// 	band.Data[i] = val
	// }
	return nil
}

// readValueByPixelType 根据像素类型读取值
func readValueByPixelType(r io.Reader, byteOrder binary.ByteOrder, pixelType int) (float64, error) {
	switch pixelType {
	case 0, 1, 2, 4: // 布尔或各种无符号整数
		var val uint8
		if err := binary.Read(r, byteOrder, &val); err != nil {
			return 0, err
		}
		return float64(val), nil
	case 3: // 有符号8位整数
		var val int8
		if err := binary.Read(r, byteOrder, &val); err != nil {
			return 0, err
		}
		return float64(val), nil
	case 5: // 有符号16位整数
		var val int16
		if err := binary.Read(r, byteOrder, &val); err != nil {
			return 0, err
		}
		return float64(val), nil
	case 6: // 无符号16位整数
		var val uint16
		if err := binary.Read(r, byteOrder, &val); err != nil {
			return 0, err
		}
		return float64(val), nil
	case 7: // 有符号32位整数
		var val int32
		if err := binary.Read(r, byteOrder, &val); err != nil {
			return 0, err
		}
		return float64(val), nil
	case 8: // 无符号32位整数
		var val uint32
		if err := binary.Read(r, byteOrder, &val); err != nil {
			return 0, err
		}
		return float64(val), nil
	case 10: // 32位浮点数
		var val float32
		if err := binary.Read(r, byteOrder, &val); err != nil {
			return 0, err
		}
		return float64(val), nil
	case 11: // 64位浮点数
		var val float64
		if err := binary.Read(r, byteOrder, &val); err != nil {
			return 0, err
		}
		return val, nil
	default:
		return 0, fmt.Errorf("不支持的像素类型: %d", pixelType)
	}
}

// readValuesByPixelType
func readValuesByPixelType(r io.Reader, byteOrder binary.ByteOrder, pixelType int, count int) ([]float64, error) {
	switch pixelType {
	case 0, 1, 2, 4: // 布尔或各种无符号整数
		data := make([]uint8, count)
		if err := binary.Read(r, byteOrder, &data); err != nil {
			return nil, fmt.Errorf("读取像素失败: %v", err)
		}
		return readAsFloat64Slice(data), nil
	case 3: // 有符号8位整数
		data := make([]int8, count)
		if err := binary.Read(r, byteOrder, &data); err != nil {
			return nil, fmt.Errorf("读取像素失败: %v", err)
		}
		return readAsFloat64Slice(data), nil
	case 5: // 有符号16位整数
		data := make([]int16, count)
		if err := binary.Read(r, byteOrder, &data); err != nil {
			return nil, fmt.Errorf("读取像素失败: %v", err)
		}
		return readAsFloat64Slice(data), nil
	case 6: // 无符号16位整数
		data := make([]uint16, count)
		if err := binary.Read(r, byteOrder, &data); err != nil {
			return nil, fmt.Errorf("读取像素失败: %v", err)
		}
		return readAsFloat64Slice(data), nil
	case 7: // 有符号32位整数
		data := make([]int32, count)
		if err := binary.Read(r, byteOrder, &data); err != nil {
			return nil, fmt.Errorf("读取像素失败: %v", err)
		}
		return readAsFloat64Slice(data), nil
	case 8: // 无符号32位整数
		data := make([]uint32, count)
		if err := binary.Read(r, byteOrder, &data); err != nil {
			return nil, fmt.Errorf("读取像素失败: %v", err)
		}
		return readAsFloat64Slice(data), nil
	case 10: // 32位浮点数
		data := make([]float32, count)
		if err := binary.Read(r, byteOrder, &data); err != nil {
			return nil, fmt.Errorf("读取像素失败: %v", err)
		}
		return readAsFloat64Slice(data), nil
	case 11: // 64位浮点数
		data := make([]float64, count)
		if err := binary.Read(r, byteOrder, &data); err != nil {
			return nil, fmt.Errorf("读取像素失败: %v", err)
		}
		return data, nil
	default:
		return nil, fmt.Errorf("不支持的像素类型")
	}
}

// WriteWKBRaster 将栅格数据写入WKB格式
func WriteWKBRaster(w io.Writer, raster *WKBRaster) error {
	// 写入字节序标识（使用小端序）
	endianByte := byte(1) // NDR/小端序
	if err := binary.Write(w, binary.LittleEndian, endianByte); err != nil {
		return fmt.Errorf("写入字节序失败: %v", err)
	}

	// 验证栅格数据
	if err := validateRaster(raster); err != nil {
		return fmt.Errorf("栅格数据验证失败: %v", err)
	}

	// 一次性写入所有头信息
	if err := binary.Write(w, binary.LittleEndian, raster.Header); err != nil {
		return fmt.Errorf("写入栅格头信息失败: %v", err)
	}

	// 写入所有波段
	for i := uint16(0); i < raster.Header.Bands; i++ {
		if i >= uint16(len(raster.Bands)) {
			return fmt.Errorf("波段数量不匹配: 头信息中指定%d个波段, 实际只有%d个", raster.Header.Bands, len(raster.Bands))
		}

		if err := writeBand(w, raster.Bands[i], raster.Header.Width, raster.Header.Height); err != nil {
			return fmt.Errorf("写入波段%d失败: %v", i, err)
		}
	}

	return nil
}

// validateRaster 验证栅格数据的有效性
func validateRaster(raster *WKBRaster) error {
	// 验证头信息
	if err := validateHeader(&raster.Header); err != nil {
		return err
	}

	// 验证波段数量匹配
	if raster.Header.Bands != uint16(len(raster.Bands)) {
		return fmt.Errorf("波段数量不匹配: 头信息中指定%d个波段, 实际有%d个", raster.Header.Bands, len(raster.Bands))
	}

	// 验证每个波段
	for i, band := range raster.Bands {
		if err := validateBand(band, raster.Header.Width, raster.Header.Height); err != nil {
			return fmt.Errorf("波段%d验证失败: %v", i, err)
		}
	}

	return nil
}

// validateBand 验证波段数据的有效性
func validateBand(band Band, width, height uint16) error {
	// 检查像素类型是否有效
	if band.Header.PixelType < 0 || band.Header.PixelType > 10 {
		return fmt.Errorf("无效的像素类型: %d", band.Header.PixelType)
	}

	// 如果不是离线数据，检查像素数据
	if !band.Header.IsOffline {
		totalPixels := int(width) * int(height)
		if len(band.Data) != totalPixels {
			return fmt.Errorf("像素数据数量不匹配: 期望%d, 实际%d", totalPixels, len(band.Data))
		}
	}

	return nil
}

// writeBand 写入单个波段数据
func writeBand(w io.Writer, band Band, width, height uint16) error {
	// 构建波段头字节
	headerByte := buildBandHeader(band.Header)

	// 写入波段头
	if err := binary.Write(w, binary.LittleEndian, headerByte); err != nil {
		return fmt.Errorf("写入波段头失败: %v", err)
	}

	// 写入无数据值
	if err := writeValueByPixelType(w, binary.LittleEndian, band.Header.PixelType, band.Header.NoData); err != nil {
		return fmt.Errorf("写入无数据值失败: %v", err)
	}

	if band.Header.IsOffline {
		// 写入离线数据
		if err := writeOfflineBand(w, band); err != nil {
			return err
		}
	} else {
		// 写入像素数据
		if err := writePixelData(w, binary.LittleEndian, band.Header.PixelType, band.Data); err != nil {
			return fmt.Errorf("写入像素数据失败: %v", err)
		}
	}

	return nil
}

// writeOfflineBand 写入离线波段数据
func writeOfflineBand(w io.Writer, band Band) error {
	// 写入波段编号（0-based）
	bandNum := byte(0) // 默认使用第一个波段
	if err := binary.Write(w, binary.LittleEndian, bandNum); err != nil {
		return fmt.Errorf("写入波段编号失败: %v", err)
	}

	// 写入以null结尾的文件路径
	if _, err := w.Write([]byte(band.Path)); err != nil {
		return fmt.Errorf("写入文件路径失败: %v", err)
	}
	if err := binary.Write(w, binary.LittleEndian, byte(0)); err != nil {
		return fmt.Errorf("写入路径终止符失败: %v", err)
	}

	return nil
}

// writePixelData 写入像素数据
func writePixelData(w io.Writer, byteOrder binary.ByteOrder, pixelType int, data []float64) error {
	for i, val := range data {
		if err := writeValueByPixelType(w, byteOrder, pixelType, val); err != nil {
			return fmt.Errorf("写入像素%d失败: %v", i, err)
		}
	}
	return nil
}

// writeValueByPixelType 根据像素类型写入值
func writeValueByPixelType(w io.Writer, byteOrder binary.ByteOrder, pixelType int, value float64) error {
	switch pixelType {
	case 0, 1, 2, 4: // 布尔或各种无符号整数
		if value < 0 || value > 255 {
			return fmt.Errorf("无符号8位整数超出范围: %v", value)
		}
		return binary.Write(w, byteOrder, uint8(value))
	case 3: // 有符号8位整数
		if value < -128 || value > 127 {
			return fmt.Errorf("有符号8位整数超出范围: %v", value)
		}
		return binary.Write(w, byteOrder, int8(value))
	case 6: // 无符号16位整数
		if value < 0 || value > 65535 {
			return fmt.Errorf("无符号16位整数超出范围: %v", value)
		}
		return binary.Write(w, byteOrder, uint16(value))
	case 5: // 有符号16位整数
		if value < -32768 || value > 32767 {
			return fmt.Errorf("有符号16位整数超出范围: %v", value)
		}
		return binary.Write(w, byteOrder, int16(value))
	case 8: // 无符号32位整数
		if value < 0 || value > 4294967295 {
			return fmt.Errorf("无符号32位整数超出范围: %v", value)
		}
		return binary.Write(w, byteOrder, uint32(value))
	case 7: // 有符号32位整数
		if value < -2147483648 || value > 2147483647 {
			return fmt.Errorf("有符号32位整数超出范围: %v", value)
		}
		return binary.Write(w, byteOrder, int32(value))
	case 9: // 32位浮点数
		return binary.Write(w, byteOrder, float32(value))
	case 10: // 64位浮点数
		return binary.Write(w, byteOrder, value)
	default:
		return fmt.Errorf("不支持的像素类型: %d", pixelType)
	}
}
