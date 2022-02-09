package ds_test

import (
	"os"
	"strings"
	"testing"

	"github.com/mocheer/pluto/pkg/assert"
	"github.com/mocheer/pluto/pkg/fn"
)

func TestRead(t *testing.T) {
	data, err := os.ReadFile("./a_test.go")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, strings.HasPrefix(fn.B2S(data), "package ds_test"), true)
}
