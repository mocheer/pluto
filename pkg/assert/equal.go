package assert

import "testing"

type Test struct {
	*testing.T
}

func New(t *testing.T) *Test {
	return &Test{t}
}

func (t Test) Equal(a, b interface{}) {
	Equal(t.T, a, b)
}

func Equal(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("Not Equal. %d %d", a, b)
	}
}
