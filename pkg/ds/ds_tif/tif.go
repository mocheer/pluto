package ds_tif

import (
	"bytes"

	"compress/zlib"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"

	"github.com/google/tiff"
	"github.com/google/tiff/bigtiff"
	_ "github.com/google/tiff/geotiff"

	"github.com/mocheer/pluto/pkg/ds/ds_tif/lzw"
)

type DsTif struct {
	Tif       tiff.TIFF
	Data      []float64
	IFD_Index int
}

// HasField 判断是否存在tagID对那个的field
func (m DsTif) HasField(tagID uint16) bool {
	ifds := m.Tif.IFDs()
	index := m.IFD_Index
	for index >= 0 {
		if ifds[index].HasField(tagID) {
			return true
		}
		index--
	}
	return false
	// return m.Tif.IFDs()[m.IFD_Index].HasField(tagID)
}

// GetField 根据tagID获取对应field
func (m DsTif) GetField(tagID uint16) tiff.Field {
	ifds := m.Tif.IFDs()
	index := m.IFD_Index
	for index >= 0 {
		if ifds[index].HasField(tagID) {
			return ifds[index].GetField(tagID)
		}
		index--
	}
	return nil
	// return m.Tif.IFDs()[m.IFD_Index].GetField(tagID)
}

// GetFirstInt
func (m DsTif) GetFirstInt(tagID uint16) uint {
	field := m.GetField(tagID)
	if field == nil {
		fmt.Printf("没有tagID:%d", tagID)
		return 0
	}
	val := field.Value()
	bs := val.Bytes()
	size := field.Type().Size()
	//
	switch field.Type() {
	case tiff.FTByte:
		return uint(bs[0])
	case tiff.FTShort: //size不一定等于2，特别是bigtiff数据格式
		return uint(val.Order().Uint16(bs[:size]))
	case tiff.FTLong: //
		return uint(val.Order().Uint32(bs[:size]))
	}
	return 0
}

// GetInt
func (m DsTif) GetInt(tagID uint16) []uint {
	field := m.GetField(tagID)
	count := field.Count()
	size := field.Type().Size()
	u := make([]uint, count)
	val := field.Value()
	bs := val.Bytes()

	switch field.Type() {
	case tiff.FTByte:
		for i := uint64(0); i < count; i++ {
			u[i] = uint(bs[i])
		}
	case tiff.FTShort:
		for i := uint64(0); i < count; i++ {
			u[i] = uint(val.Order().Uint16(bs[size*i : size*(i+1)]))
		}
	case tiff.FTLong:
		for i := uint64(0); i < count; i++ {
			u[i] = uint(val.Order().Uint32(bs[size*i : size*(i+1)]))
		}
	case bigtiff.FTLong8:
		for i := uint64(0); i < count; i++ {
			u[i] = uint(val.Order().Uint32(bs[size*i : size*(i+1)]))
		}
	}
	return u
}

// GetFloat
func (m DsTif) GetFloat(tagID uint16) []float64 {
	field := m.GetField(tagID)
	count := field.Count()
	u := make([]float64, count)
	val := field.Value()
	bs := val.Bytes()
	switch field.Type() {
	case tiff.FTFloat:
		u2 := make([]float32, count)
		for i := uint64(0); i < count; i++ {
			// I'm not sure this code will work
			buf := bytes.NewReader(bs[4*i : 4*(i+1)])
			binary.Read(buf, val.Order(), &u2[i])
		}
		for i := uint64(0); i < count; i++ {
			u[i] = float64(u2[i])
		}
	case tiff.FTDouble:
		for i := uint64(0); i < count; i++ {
			buf := bytes.NewReader(bs[8*i : 8*(i+1)])
			binary.Read(buf, val.Order(), &u[i])
		}

	}
	return u
}

// Origin 获取坐标原点，通常是经纬度坐标原点
func (m DsTif) Origin() []float64 {
	origin := m.GetFloat(tModelTiepointTag)
	return origin[3:6]
}

// 栅格坐标和模型坐标的比例
func (m DsTif) Scale() []float64 {
	return m.GetFloat(tModelPixelScaleTag)
}

// 栅格坐标到模型坐标的变换矩阵，这个值很多时候是空值
func (m DsTif) Transform() []float64 {
	return m.GetFloat(tModelTransformationTag)
}

// GetColRow 获取经纬度对应的行列坐标，这种算法不一定精确，不一定等间隔
func (m DsTif) GetColRow(lon, lat float64) [2]int {
	origin := m.Origin()
	scale := m.Scale()
	col := int(math.Round((lon - origin[0]) / scale[0]))
	row := int(math.Round((origin[1] - lat) / scale[1]))
	width, height := m.Width(), m.Height()
	if row >= height {
		row = height - 1
	}
	if col >= width {
		col = width - 1
	}
	return [2]int{col, row}
}

// GetLonLat 获取行列坐标对应的经纬度
func (m DsTif) GetLonLat(col, row int) [2]float64 {
	ModelTiepoint := m.GetFloat(tModelTiepointTag)
	scale := m.Scale()
	//
	lon := (float64(col)-ModelTiepoint[0])*scale[0] + ModelTiepoint[3]
	lat := (float64(row)-ModelTiepoint[1])*scale[1]*(-1) + ModelTiepoint[4]

	return [2]float64{lon, lat}
}

// BBox
func (m DsTif) BBox() [4]float64 {
	origin := m.Origin()
	scale := m.Scale()
	var maxX = origin[0] + (scale[0] * float64(m.Width()))
	var minY = origin[1] - (scale[1] * float64(m.Height()))
	return [4]float64{
		origin[0], origin[1], maxX, minY,
	}
}

// Width
func (m DsTif) Width() int {
	return int(m.GetFirstInt(tImageWidth))
}

// Height
func (m DsTif) Height() int {
	return int(m.GetFirstInt(tImageLength))
}

// GetAlt 获取行列坐标对应的高程数据
func (m DsTif) GetAlt(column, row int) float64 {
	return m.Data[column+row*m.Width()]
}

func (m DsTif) GetAltByLonLat(lon, lat float64) float64 {
	colrow := m.GetColRow(lon, lat)
	col, row := colrow[0], colrow[1]
	return m.GetAlt(col, row)
}

func (m DsTif) readData() (data []float64, err error) {
	tif := m.Tif
	compressionType := m.GetFirstInt(tCompression)
	SampleFormat := m.GetFirstInt(tSampleFormat)

	width := m.Width()
	height := m.Height()

	//
	data = make([]float64, width*height)

	var off int
	bitsPerSample := m.GetInt(tBitsPerSample)
	blockPadding := false
	blockWidth := int(width)
	blockHeight := int(height)
	blocksAcross := 1
	blocksDown := 1

	var blockOffsets, blockCounts []uint

	if m.HasField(tTileWidth) {
		tileWidth := int(m.GetFirstInt(tTileWidth))
		tileHeight := int(m.GetFirstInt(tTileLength))

		blockPadding = true

		blockWidth = int(tileWidth)
		blockHeight = int(tileHeight)

		blocksAcross = (width + blockWidth - 1) / blockWidth
		blocksDown = (height + blockHeight - 1) / blockHeight

		if ok := m.HasField(tTileOffsets); ok {
			blockOffsets = m.GetInt(tTileOffsets)
		}

		if ok := m.HasField(tTileByteCounts); ok {
			blockCounts = m.GetInt(tTileByteCounts)
		}

	} else {
		if m.HasField(tRowsPerStrip) {
			blockHeight = int(m.GetFirstInt(tRowsPerStrip))
		}

		blocksDown = (height + blockHeight - 1) / blockHeight

		if ok := m.HasField(tStripOffsets); ok {
			blockOffsets = m.GetInt(tStripOffsets)
		}

		if ok := m.HasField(tStripByteCounts); ok {
			blockCounts = m.GetInt(tStripByteCounts)
		}
	}
	var buf []byte

	for i := 0; i < blocksAcross; i++ {
		blkW := blockWidth
		if !blockPadding && i == blocksAcross-1 && width%blockWidth != 0 {
			blkW = width % blockWidth
		}
		for j := 0; j < blocksDown; j++ {
			blkH := blockHeight
			if !blockPadding && j == blocksDown-1 && height%blockHeight != 0 {
				blkH = height % blockHeight
			}
			offset := int64(blockOffsets[j*blocksAcross+i])
			n := int64(blockCounts[j*blocksAcross+i])
			switch compressionType {
			case cNone:
				buf = make([]byte, n)
				_, err = tif.R().ReadAt(buf, offset)

			case cLZW: // lzw 压缩
				// tiff文件采用类似compress/lzw但不兼容的算法
				// 目前有部分文件解析失败  unexpected EOF
				r := lzw.NewReader(io.NewSectionReader(tif.R(), offset, n), lzw.MSB, 8)
				buf, err = io.ReadAll(r)
				r.Close()
				// 这里其实可能存在错误（unexpected EOF），其实GeoTiffjs一样有这个错误，但仍可获取数据
				// 目前获取到的数据和GeoTiff存在差别
				if err != nil {
					fmt.Println(err)
					// return nil, err
				}
			case cDeflate, cDeflateOld:
				r, err := zlib.NewReader(io.NewSectionReader(tif.R(), offset, n))
				if err != nil {
					return nil, err
				}
				buf, err = io.ReadAll(r)
				r.Close()
				if err != nil {
					return nil, err
				}
			case cPackBits:

			default:
				err = fmt.Errorf("unsupported compression value %d", compressionType)
			}

			xmin := i * blockWidth
			ymin := j * blockHeight
			xmax := xmin + blkW
			ymax := ymin + blkH

			xmax = minInt(xmax, width)
			ymax = minInt(ymax, height)

			off = 0

			// Apply horizontal predictor if necessary.
			// In this case, p contains the color difference to the preceding pixel.
			// See page 64-65 of the spec.

			if m.HasField(tPredictor) && m.GetFirstInt(tPredictor) == prHorizontal {
				// does it make sense to extend this to 32 and 64 bits?
				if bitsPerSample[0] == 16 {
					var off int
					spp := len(bitsPerSample) // samples per pixel
					bpp := spp * 2            // bytes per pixel
					for y := ymin; y < ymax; y++ {
						off += spp * 2
						for x := 0; x < (xmax-xmin-1)*bpp; x += 2 {
							v0 := tif.R().ByteOrder().Uint16(buf[off-bpp : off-bpp+2])
							v1 := tif.R().ByteOrder().Uint16(buf[off : off+2])
							tif.R().ByteOrder().PutUint16(buf[off:off+2], v1+v0)
							off += 2
						}
					}
				} else if bitsPerSample[0] == 8 {
					var off int
					spp := len(bitsPerSample) // samples per pixel
					for y := ymin; y < ymax; y++ {
						off += spp
						for x := 0; x < (xmax-xmin-1)*spp; x++ {
							buf[off] += buf[off-spp]
							off++
						}
					}
				}
			}
			var mode imageMode
			PhotometricInterp := m.GetFirstInt(tPhotometricInterpretation)
			var palette []uint32
			// Determine the image mode.
			switch PhotometricInterp {
			case PI_RGB:
				if bitsPerSample[0] == 16 {
					for _, b := range bitsPerSample {
						if b != 16 {
							err = errors.New("wrong number of samples for 16bit RGB")
							return
						}
					}
				} else {
					for _, b := range bitsPerSample {
						if b != 8 {
							err = errors.New("wrong number of samples for 8bit RGB")
							return
						}
					}
				}
				// RGB images normally have 3 samples per pixel.
				// If there are more, ExtraSamples (p. 31-32 of the spec)
				// gives their meaning (usually an alpha channel).
				//
				// This implementation does not support extra samples
				// of an unspecified type.
				switch len(bitsPerSample) {
				case 3:
					mode = mRGB
				case 4:
					switch m.GetFirstInt(tExtraSamples) {
					case 1:
						mode = mRGBA
					case 2:
						mode = mNRGBA
					default:
						err = errors.New("wrong number of samples for RGB")
						return
					}
				default:
					err = errors.New("wrong number of samples for RGB")
					return
				}
			case PI_Paletted:
				mode = mPaletted
				// retreive the palette colour data
				if ok := m.HasField(tColorMap); ok {
					val := m.GetInt(tColorMap)
					numcolors := len(val) / 3
					if len(val)%3 != 0 || numcolors <= 0 || numcolors > 256 {
						return nil, errors.New("bad ColorMap length")
					}
					palette = make([]uint32, numcolors)
					for i := 0; i < numcolors; i++ {
						// colours in the colour map are given in 16-bit channels
						// and need to be rescaled to an 8-bit format.
						red := uint32(float64(val[i]) / 65535.0 * 255.0)
						green := uint32(float64(val[i+numcolors]) / 65535.0 * 255.0)
						blue := uint32(float64(val[i+2*numcolors]) / 65535.0 * 255.0)
						a := uint32(255)
						val := uint32((a << 24) | (red << 16) | (green << 8) | blue)
						palette[i] = val
					}
				} else {
					err = errors.New("could not locate the colour map tag")
					return
				}
			case PI_WhiteIsZero:
				mode = mGrayInvert
			case PI_BlackIsZero:
				mode = mGray
			default:
				err = errors.New("unsupported image format")
				return
			}

			switch mode {
			case mGray, mGrayInvert:
				switch SampleFormat {
				case 1: // Unsigned integer data
					switch bitsPerSample[0] {
					case 8:
						for y := ymin; y < ymax; y++ {
							for x := xmin; x < xmax; x++ {
								i := y*width + x
								data[i] = float64(buf[off])
								off++
							}
						}
					case 16:
						for y := ymin; y < ymax; y++ {
							for x := xmin; x < xmax; x++ {
								value := tif.R().ByteOrder().Uint16(buf[off : off+2])
								i := y*width + x
								data[i] = float64(value)
								off += 2
							}
						}
					case 32:
						for y := ymin; y < ymax; y++ {
							for x := xmin; x < xmax; x++ {
								value := tif.R().ByteOrder().Uint32(buf[off : off+4])
								i := y*width + x
								data[i] = float64(value)
								off += 4
							}
						}
					case 64:
						for y := ymin; y < ymax; y++ {
							for x := xmin; x < xmax; x++ {
								value := tif.R().ByteOrder().Uint64(buf[off : off+8])
								i := y*width + x
								data[i] = float64(value)
								off += 8
							}
						}
					default:
						err = errors.New("unsupported data format")
						return
					}
				case 2: // Signed integer data
					switch bitsPerSample[0] {
					case 8:
						for y := ymin; y < ymax; y++ {
							for x := xmin; x < xmax; x++ {
								i := y*width + x
								data[i] = float64(int8(buf[off]))
								off++
							}
						}
					case 16:
						for y := ymin; y < ymax; y++ {
							for x := xmin; x < xmax; x++ {
								value := int16(tif.R().ByteOrder().Uint16(buf[off : off+2]))
								i := y*width + x
								data[i] = float64(value)
								off += 2
							}
						}
					case 32:
						for y := ymin; y < ymax; y++ {
							for x := xmin; x < xmax; x++ {
								value := int32(tif.R().ByteOrder().Uint32(buf[off : off+4]))
								i := y*width + x
								data[i] = float64(value)
								off += 4
							}
						}
					case 64:
						for y := ymin; y < ymax; y++ {
							for x := xmin; x < xmax; x++ {
								value := int64(tif.R().ByteOrder().Uint64(buf[off : off+8]))
								i := y*width + x
								data[i] = float64(value)
								off += 8
							}
						}
					default:
						err = errors.New("unsupported data format")
						return
					}
				case 3: // Floating point data
					switch bitsPerSample[0] {
					case 32:
						for y := ymin; y < ymax; y++ {
							for x := xmin; x < xmax; x++ {
								if off <= len(buf) {
									bits := tif.R().ByteOrder().Uint32(buf[off : off+4])
									float := math.Float32frombits(bits)
									i := y*width + x
									data[i] = float64(float)
									off += 4
								}
							}
						}
					case 64:
						for y := ymin; y < ymax; y++ {
							for x := xmin; x < xmax; x++ {
								if off <= len(buf) {
									bits := tif.R().ByteOrder().Uint64(buf[off : off+8])
									float := math.Float64frombits(bits)
									i := y*width + x
									data[i] = float
									off += 8
								}
							}
						}
					default:
						err = errors.New("unsupported data format")
						return
					}
				default:
					err = errors.New("unsupported sample format")
					return
				}
			case mPaletted:
				for y := ymin; y < ymax; y++ {
					for x := xmin; x < xmax; x++ {
						i := y*width + x
						val := int(buf[off])
						data[i] = float64(palette[val])
						off++
					}
				}

			case mRGB:
				if bitsPerSample[0] == 8 {
					for y := ymin; y < ymax; y++ {
						for x := xmin; x < xmax; x++ {
							red := uint32(buf[off])
							green := uint32(buf[off+1])
							blue := uint32(buf[off+2])
							a := uint32(255)
							off += 3
							i := y*width + x
							val := uint32((a << 24) | (red << 16) | (green << 8) | blue)
							data[i] = float64(val)
						}
					}
				} else if bitsPerSample[0] == 16 {
					for y := ymin; y < ymax; y++ {
						for x := xmin; x < xmax; x++ {
							// the spec doesn't talk about 16-bit RGB images so
							// I'm not sure why I bother with this. They specifically
							// say that RGB images are 8-bits per channel. Anyhow,
							// I rescale the 16-bits to an 8-bit channel for simplicity.
							red := uint32(float64(tif.R().ByteOrder().Uint16(buf[off+0:off+2])) / 65535.0 * 255.0)
							green := uint32(float64(tif.R().ByteOrder().Uint16(buf[off+2:off+4])) / 65535.0 * 255.0)
							blue := uint32(float64(tif.R().ByteOrder().Uint16(buf[off+4:off+6])) / 65535.0 * 255.0)
							a := uint32(255)
							off += 6
							i := y*width + x
							val := uint32((a << 24) | (red << 16) | (green << 8) | blue)
							data[i] = float64(val)
						}
					}
				} else {
					err = errors.New("unsupported data format")
					return
				}
			case mNRGBA:
				if bitsPerSample[0] == 8 {
					for y := ymin; y < ymax; y++ {
						for x := xmin; x < xmax; x++ {
							red := uint32(buf[off])
							green := uint32(buf[off+1])
							blue := uint32(buf[off+2])
							a := uint32(buf[off+3])
							off += 4
							i := y*width + x
							val := uint32((a << 24) | (red << 16) | (green << 8) | blue)
							data[i] = float64(val)
						}
					}
				} else if bitsPerSample[0] == 16 {
					for y := ymin; y < ymax; y++ {
						for x := xmin; x < xmax; x++ {
							red := uint32(float64(tif.R().ByteOrder().Uint16(buf[off+0:off+2])) / 65535.0 * 255.0)
							green := uint32(float64(tif.R().ByteOrder().Uint16(buf[off+2:off+4])) / 65535.0 * 255.0)
							blue := uint32(float64(tif.R().ByteOrder().Uint16(buf[off+4:off+6])) / 65535.0 * 255.0)
							a := uint32(float64(tif.R().ByteOrder().Uint16(buf[off+6:off+8])) / 65535.0 * 255.0)
							off += 8
							i := y*width + x
							val := uint32((a << 24) | (red << 16) | (green << 8) | blue)
							data[i] = float64(val)
						}
					}
				} else {
					err = errors.New("unsupported data format")
					return
				}
			case mRGBA:
				if bitsPerSample[0] == 16 {
					for y := ymin; y < ymax; y++ {
						for x := xmin; x < xmax; x++ {
							red := uint32(float64(tif.R().ByteOrder().Uint16(buf[off+0:off+2])) / 65535.0 * 255.0)
							green := uint32(float64(tif.R().ByteOrder().Uint16(buf[off+2:off+4])) / 65535.0 * 255.0)
							blue := uint32(float64(tif.R().ByteOrder().Uint16(buf[off+4:off+6])) / 65535.0 * 255.0)
							a := uint32(float64(tif.R().ByteOrder().Uint16(buf[off+6:off+8])) / 65535.0 * 255.0)
							off += 8
							i := y*width + x
							val := uint32((a << 24) | (red << 16) | (green << 8) | blue)
							data[i] = float64(val)
						}
					}
				} else {
					for y := ymin; y < ymax; y++ {
						for x := xmin; x < xmax; x++ {
							red := uint32(buf[off])
							green := uint32(buf[off+1])
							blue := uint32(buf[off+2])
							a := uint32(buf[off+3])
							off += 4
							i := y*width + x
							val := uint32((a << 24) | (red << 16) | (green << 8) | blue)
							data[i] = float64(val)
						}
					}
				}
			}
		}
	}

	return
}
