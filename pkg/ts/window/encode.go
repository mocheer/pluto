package window

import (
	"bytes"
	"fmt"
)

// encodeURI 与JavaScript encodeURI()完全等价
func EncodeURI(uri string) string {
	var buf bytes.Buffer
	buf.Grow(len(uri) * 3) // 预分配空间

	for i := 0; i < len(uri); {
		b := uri[i]

		// ASCII字符处理
		if b < 0x80 {
			if shouldNotEncodeASCII(b) {
				buf.WriteByte(b)
			} else {
				// 对需要编码的ASCII字符进行百分号编码
				fmt.Fprintf(&buf, "%%%02X", b)
			}
			i++
			continue
		}

		// 多字节UTF-8字符处理
		r, size := DecodeUTF8Char(uri[i:])
		if r == 0 {
			// 无效UTF-8，按原样保留字节
			fmt.Fprintf(&buf, "%%%02X", b)
			i++
		} else {
			// 对非ASCII字符的UTF-8字节进行编码
			for j := 0; j < size; j++ {
				fmt.Fprintf(&buf, "%%%02X", uri[i+j])
			}
			i += size
		}
	}

	return buf.String()
}

// decodeUTF8Char 解码UTF-8字符并返回rune和字节数
func DecodeUTF8Char(s string) (rune, int) {
	if len(s) == 0 {
		return 0, 0
	}

	b1 := s[0]
	switch {
	case b1 < 0x80:
		return rune(b1), 1
	case b1 < 0xE0:
		if len(s) < 2 {
			return 0, 0
		}
		b2 := s[1]
		if b2&0xC0 != 0x80 {
			return 0, 0
		}
		r := rune(b1&0x1F)<<6 | rune(b2&0x3F)
		if r < 0x80 {
			return 0, 0 // 非最小化编码
		}
		return r, 2
	case b1 < 0xF0:
		if len(s) < 3 {
			return 0, 0
		}
		b2, b3 := s[1], s[2]
		if b2&0xC0 != 0x80 || b3&0xC0 != 0x80 {
			return 0, 0
		}
		r := rune(b1&0x0F)<<12 | rune(b2&0x3F)<<6 | rune(b3&0x3F)
		if r < 0x800 {
			return 0, 0 // 非最小化编码
		}
		return r, 3
	default:
		if len(s) < 4 {
			return 0, 0
		}
		b2, b3, b4 := s[1], s[2], s[3]
		if b2&0xC0 != 0x80 || b3&0xC0 != 0x80 || b4&0xC0 != 0x80 {
			return 0, 0
		}
		r := rune(b1&0x07)<<18 | rune(b2&0x3F)<<12 | rune(b3&0x3F)<<6 | rune(b4&0x3F)
		if r < 0x10000 {
			return 0, 0 // 非最小化编码
		}
		return r, 4
	}
}

// shouldNotEncodeASCII 判断ASCII字符是否需要编码
func shouldNotEncodeASCII(b byte) bool {
	// RFC 3986不编码的ASCII字符
	// 字母数字
	if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
		return true
	}

	// 标记
	switch b {
	case '-', '_', '.', '!', '~', '*', '\'', '(', ')':
		return true
	}

	// 保留字符
	switch b {
	case ';', ',', '/', '?', ':', '@', '&', '=', '+', '$', '#':
		return true
	}

	// 所有其他ASCII字符都需要编码
	return false
}
