package fn_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/fn"
	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	var result string
	result = fn.Format("{a}", map[string]any{"a": 1})
	assert.Equal(t, result, "1")
	type Model struct {
		A string
	}
	result = fn.Format("{a}", Model{A: "1"})
	assert.Equal(t, result, "1")
}

func BenchmarkFormat(b *testing.B) {
	fn.Format("{a}", map[string]any{"a": 1})
}
