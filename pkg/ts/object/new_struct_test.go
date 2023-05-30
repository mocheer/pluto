package object_test

import (
	"fmt"
	"testing"

	"github.com/mocheer/pluto/pkg/ts/object"
)

func TestNewStruct(t *testing.T) {
	pe := object.NewStructBuilder().
		AddString("Name", "").
		AddInt64("Age", "").
		AddFloat64("Height", "").
		AddStringSlice("Scores", "").
		AddFunc("MyFunction", "").
		Build()

	p := pe.New()
	p.SetString("Name", "你好")
	p.SetInt64("Age", 32)
	p.SetFloat64("Height", 100)
	p.SetStringSlice("Scores", []string{"1", "2"})
	p.SetFunc("MyFunction", func() { fmt.Println("my function") })

	fmt.Printf("%+v\n", p)
	fmt.Printf("%T，%+v\n", p.Interface(), p.Interface())
	fmt.Printf("%T，%+v\n", p.Addr(), p.Addr())

}
