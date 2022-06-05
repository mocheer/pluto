package assert

import "testing"

type Test struct {
	Ctx testing.TB
}

func New(t testing.TB) *Test {
	return &Test{t}
}

//
func (m Test) Equal(a, b interface{}) {
	if a != b {
		m.Ctx.Errorf("Not Equal. %d %d", a, b)
	}
}

//
func (m Test) NotEqual(a, b interface{}) {
	if a == b {
		m.Ctx.Errorf("Equal. %d %d", a, b)
	}
}
