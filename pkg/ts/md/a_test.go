package md_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/ts/md"
)

func TestMH(t *testing.T) {
	data := md.MD([]byte("## markdown document   ")).HTML()
	t.Log(data)
}

// func TestMH2(t *testing.T) {
// 	data := md.MD([]byte("## markdown document   ")).HTML2()
// 	t.Log(data)
// }
