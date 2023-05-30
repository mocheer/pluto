package ds_toml_test

import (
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
	ds_toml.ReadFile("./testdata/a_test.toml", &conf)
	t.Log(conf)
}

func TestSave(t *testing.T) {
	var conf = &Config{}
	conf.Age = 30
	err := ds_toml.Save("./testdata/a_test2.toml", conf)
	t.Log(err)
}
