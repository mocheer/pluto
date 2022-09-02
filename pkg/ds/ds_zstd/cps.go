package ds_zstd

import (
	"bytes"

	"github.com/klauspost/compress/zstd"
)

// Encode input to output.
func Encode(data []byte) ([]byte, error) {
	var b bytes.Buffer
	enc, err := zstd.NewWriter(&b)
	if err != nil {
		return nil, err
	}
	defer enc.Close()
	enc.Write(data)
	return b.Bytes(), err
}
