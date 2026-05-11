package ds_docx

import (
	"archive/zip"
	"io"
	"os"
)

// ReadFile
func ReadFile(fname string) (*WordprocessingML, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	// defer f.Close()
	size, err := f.Stat()
	if err != nil {
		return nil, err
	}
	return Read(f, size.Size())
}

// Read
func Read(r io.ReaderAt, size int64) (*WordprocessingML, error) {
	z, err := zip.NewReader(r, size)
	if err != nil {
		return nil, err
	}
	return &WordprocessingML{reader: z}, nil
}
