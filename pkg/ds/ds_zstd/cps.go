package ds_zstd

import (
	"bytes"
	"io"

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

// Decode
func Decode(data []byte) ([]byte, error) {
	dr := bytes.NewReader(data)
	d, err := zstd.NewReader(dr)
	if err != nil {
		return nil, err
	}
	defer d.Close()

	buf, err := io.ReadAll(d)
	if err != nil {
		return buf, err
	}
	return buf, err
}

// Compress input to output.
func Compress(in io.Reader, out io.Writer) error {
	enc, err := zstd.NewWriter(out)
	if err != nil {
		return err
	}
	_, err = io.Copy(enc, in)
	if err != nil {
		enc.Close()
		return err
	}
	return enc.Close()
}

func Decompress(in io.Reader, out io.Writer) error {
	d, err := zstd.NewReader(in)
	if err != nil {
		return err
	}
	defer d.Close()

	// Copy content...
	_, err = io.Copy(out, d)
	return err
}
