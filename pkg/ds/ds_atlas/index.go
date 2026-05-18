package atlas

import (
	"bytes"
	"image/png"
	"os"
	"path/filepath"
)

// Read 拆解指定的 atlas 合图文件，返回所有 Atlas 对象。
func Read(name string) (*Atlas, error) {
	return parse(name)
}

func (a *Atlas) Save(dir string) error {
	for name, f := range a.Frames {
		buffer := bytes.NewBuffer(nil)
		err := png.Encode(buffer, f.SubImage)
		if err != nil {
			return err
		}
		// write data to file
		if err := os.WriteFile(filepath.Join(dir, name), buffer.Bytes(), 0644); err != nil {
			return err
		}
	}
	return nil
}
