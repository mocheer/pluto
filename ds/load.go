package ds

import (
	"io"
	"net/http"
)

// Load 下载保存
func Load(url string, fileName string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	f := MustCreate(fileName)
	defer f.Close()
	io.Copy(f, res.Body)
	return nil
}
