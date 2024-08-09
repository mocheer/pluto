package ds_test

import (
	"fmt"
	"net/url"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/fn"
	"github.com/samber/lo"
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

func TestMerge(t *testing.T) {
	files, _ := ds.GetFiles("D:\\data\\apk\\从前一条街2\\assets")
	zipFiles := lo.Filter(files, func(f string, _ int) bool {
		return strings.Contains(f, ".zip")
	})
	slices.SortFunc(zipFiles, func(s1 string, s2 string) int {
		s1 = strings.Replace(strings.Replace(s1, "D:\\data\\apk\\从前一条街2\\assets\\data", "", -1), ".zip", "", -1)
		s2 = strings.Replace(strings.Replace(s2, "D:\\data\\apk\\从前一条街2\\assets\\data", "", -1), ".zip", "", -1)
		i1, _ := strconv.Atoi(s1)
		i2, _ := strconv.Atoi(s2)
		return i1 - i2
	})
	t.Log(zipFiles)
	err := ds.MergeFiles(zipFiles, "./testdata/a.zip")
	t.Log(err)
}
