package icolor

import (
	"fmt"
	"image/color"
	"strconv"
)

// FromHex 解析16进制字符串颜色
func FromHex(hex string) color.RGBA {
	var r, g, b, a uint64
	if hex[0:1] == "#" {
		hex = hex[1:]
	}
	switch len(hex) {
	case 2:
		r, _ = strconv.ParseUint(hex, 16, 0)
		g = r
		b = r
		a = 255
	case 3:
		r, _ = strconv.ParseUint(hex[0:1]+hex[0:1], 16, 0)
		g, _ = strconv.ParseUint(hex[1:2]+hex[1:2], 16, 0)
		b, _ = strconv.ParseUint(hex[2:3]+hex[2:3], 16, 0)
		a = 255
	case 6:
		r, _ = strconv.ParseUint(hex[0:2], 16, 0)
		g, _ = strconv.ParseUint(hex[2:4], 16, 0)
		b, _ = strconv.ParseUint(hex[4:6], 16, 0)
		a = 255
	case 8:
		r, _ = strconv.ParseUint(hex[0:2], 16, 0)
		g, _ = strconv.ParseUint(hex[2:4], 16, 0)
		b, _ = strconv.ParseUint(hex[4:6], 16, 0)
		a, _ = strconv.ParseUint(hex[6:8], 16, 0)
	default:
		return color.RGBA{}
	}

	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
}

var rgbaHexFormat = "%02X%02X%02X%02X"

func ToHex(c color.Color) string {
	r, g, b, a := c.RGBA()
	return fmt.Sprintf(rgbaHexFormat, r&0xff, g&0xff, b&0xff, a&0xff)
}
