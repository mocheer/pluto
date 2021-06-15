package tileset

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mocheer/pluto/fn"
	"github.com/mocheer/pluto/fs"
)

func Load(remoteURL, dirName string) error {
	data, err := fn.Load(remoteURL)
	if err == nil {
		name := filepath.Base(remoteURL) // tileset.json
		fs.SaveFile(filepath.Join(dirName, name), data)
		//
		tiles := FromBytes(data)
		loadTile(tiles.Root, remoteURL[:strings.LastIndex(remoteURL, "/")], dirName)
	}
	return err
}

func loadTile(t Tile, baseURL, dirName string) error {
	contentURL := t.Content.Url
	if contentURL == "" {
		contentURL = t.Content.Uri
	}
	if contentURL != "" {
		contentRemoteURL := baseURL + "/" + contentURL
		relativePath := dirName + "/" + contentURL
		fmt.Println(contentRemoteURL, relativePath)
		//
		if strings.HasSuffix(contentURL, ".json") {
			Load(contentRemoteURL, filepath.Dir(relativePath))
		} else {
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
