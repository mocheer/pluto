package ds_brotli

import (
	"bytes"

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
