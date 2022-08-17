package dsfileinfo_test

import (
	"os"
	"testing"

	"github.com/floyernick/fleep-go"
)

func TestImage(t *testing.T) {
	file, _ := os.ReadFile("D:\\code-space\\go\\skybox\\nix\\data\\GaodeMap.Normal\\3\\2\\2.png") // Reads PNG file
	info, _ := fleep.GetInfo(file)                                                                // Gets file format
	t.Log(info.Type)                                                                              // Prints [raster-image]
	t.Log(info.Extension)                                                                         // Prints [png]
	t.Log(info.Mime)                                                                              // Prints [image/png]
	t.Log(info.TypeMatches(fleep.RasterImage))                                                    // Prints true
	t.Log(info.ExtensionMatches("jpg"))                                                           // Prints false
	t.Log(info.MimeMatches("image/png"))                                                          // Prints true
}
