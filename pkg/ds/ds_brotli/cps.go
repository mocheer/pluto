package ds_brotli

import (
	"bytes"
	"io"

	"github.com/andybalholm/brotli"
)

// Encode input to output.
func Encode(data []byte) []byte {
	var b bytes.Buffer
	w := brotli.NewWriter(&b)
	defer w.Close()
	w.Write(data)

	return b.Bytes()
}

// Decode
func Decode(data []byte) ([]byte, error) {
	dr := bytes.NewReader(data)
	br := brotli.NewReader(dr)
	buf, err := io.ReadAll(br)
	if err != nil {
		return buf, err
	}
	return buf, nil
}
