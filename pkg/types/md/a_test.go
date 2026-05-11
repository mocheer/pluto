package md_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/types/md"
)

func TestMH(t *testing.T) {
	m := md.New()
	m.Write([]byte("## markdown document   "))
	data := m.ToHTML()
	t.Log(data)
}

// func TestMH2(t *testing.T) {
// 	data := md.MD([]byte("## markdown document   ")).HTML2()
// 	t.Log(data)
// }
