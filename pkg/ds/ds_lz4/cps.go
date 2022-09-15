package ds_lz4

import (
	lz4 "github.com/pierrec/lz4/v4"
)

// Encode
func Encode(data []byte) ([]byte, error) {
	buf := make([]byte, lz4.CompressBlockBound(len(data)))
	var c lz4.Compressor
	n, err := c.CompressBlock(data, buf)
	if err != nil || n >= len(data) {
		return nil, err
	}
	return buf[:n], nil
}

// Decode
func Decode(buf []byte) ([]byte, error) {
	// Allocate a very large buffer for decompression.
	out := make([]byte, 10*len(buf))
	n, err := lz4.UncompressBlock(buf, out)
	if err != nil {
		return nil, nil
	}
	return out[:n], nil
}
