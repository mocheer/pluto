package icolor

import (
	"fmt"
	"image/color"
	"regexp"
	"strconv"
)

// FromHtmlColor 从网页的颜色代码获取color
func FromHtmlColor(value string) color.RGBA {
	if value[0] == '#' {
		return FromHex(value)
	}
	return FromRGBAString(value)
}

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

func FromRGBAString(rgba string) color.RGBA {

	// 定义正则表达式，提取四个数字
	reg := regexp.MustCompile(`(\d+)\W+(\d+)\W+(\d+)\W+([\d.]+)`)
	matches := reg.FindStringSubmatch(rgba)

	// 解析数字字符串
	r, _ := strconv.ParseInt(matches[1], 10, 0)
	g, _ := strconv.ParseInt(matches[2], 10, 0)
	b, _ := strconv.ParseInt(matches[3], 10, 0)
	alpha, _ := strconv.ParseFloat(matches[4], 64)

	return color.RGBA{
		uint8(r),
		uint8(g),
		uint8(b),
		uint8(alpha * 255),
	}
}

func ToHex(c color.Color) string {
	r, g, b, a := c.RGBA()
	return fmt.Sprintf("%02X%02X%02X%02X", r&0xff, g&0xff, b&0xff, a&0xff)
}
