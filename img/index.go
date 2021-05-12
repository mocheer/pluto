package img

import (
	"image"
	"os"
)

// FromFile
func FromFile(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	image, _, err := image.Decode(f)
	return image, err
}
