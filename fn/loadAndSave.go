package fn

import (
	"io"
	"net/http"
	"os"
)

// LoadAndSave 下载保存
func LoadAndSave(url string, filePath string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	io.Copy(f, res.Body)
	return nil
}
