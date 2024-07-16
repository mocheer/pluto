package ds_test

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/fn"
	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	data, err := os.ReadFile("./a_test.go")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, strings.HasPrefix(fn.BytesToString(data), "package ds_test"), true)
}

func TestRename(t *testing.T) {
	// err := ds.EachDirsToRename("D:\\code\\go\\nix\\data\\3dtiles-20240708143021", func(oldName string) string {
	// 	name, err := url.QueryUnescape(oldName)
	// 	fmt.Println(name)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	return name
	// })
	// t.Log(err)
	// ds.Each("D:\\code\\go\\nix\\data\\3dtiles-20240708143021", func(filename string, fi os.FileInfo) {
	// 	fmt.Println(url.QueryUnescape(fi.Name()))
	// })

	// dirs, _ := ds.GetDirs("D:\\code\\go\\nix\\data\\3dtiles-20240708143021")
	// for _, dir := range dirs {
	// 	name, _ := url.QueryUnescape(dir)
	// 	os.Rename(dir, name)
	// }

	ds.EachFilesToRename("D:\\code\\go\\nix\\data\\3dtiles-20240708143021", func(oldName string) string {
		name, err := url.QueryUnescape(oldName)
		fmt.Println(name)
		if err != nil {
			panic(err)
		}
		return name
	})
}
