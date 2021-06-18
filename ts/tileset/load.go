package tileset

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mocheer/pluto/fs"
	"github.com/mocheer/pluto/rh"
)

// Load 下载远程地址中的tileset以及其所有的tile
func Load(remoteURL, dirName string) error {
	data, err := rh.Get(remoteURL)
	if err == nil {
		name := filepath.Base(remoteURL) // tileset.json
		fs.Save(filepath.Join(dirName, name), data)
		//
		tiles := FromBytes(data)
		loadTile(tiles.Root, remoteURL[:strings.LastIndex(remoteURL, "/")], dirName)
	}
	return err
}

// loadTile 下载tile以及tile的所有子级
func loadTile(t Tile, baseURL, dirName string) error {
	contentURL := t.Content.Url
	if contentURL == "" {
		contentURL = t.Content.Uri
	}
	if contentURL != "" {
		contentRemoteURL := baseURL + "/" + contentURL
		relativePath := dirName + "/" + contentURL
		fmt.Println(contentRemoteURL)
		// 这是一个json描述文件
		if strings.HasSuffix(contentURL, ".json") {
			Load(contentRemoteURL, filepath.Dir(relativePath))
		} else { // 一个b3dm文件
			err := fs.Load(contentRemoteURL, relativePath)
			if err != nil {
				return err
			}
		}
	}

	children := t.Children
	if len(children) > 0 {
		for _, t := range children {
			err := loadTile(t, baseURL, dirName)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
