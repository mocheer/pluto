package ds_toml_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/mocheer/pluto/pkg/ds/ds_toml"
)

type Config struct {
	Age        int
	Cats       []string
	Pi         float64
	Perfection []int
	DOB        time.Time
}

func TestRead(t *testing.T) {
	var conf Config
	ds_toml.ReadFile("./a_test.toml", &conf)
	fmt.Println(conf)
}
