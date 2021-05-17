package assert

import "testing"

func Equal(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("Not Equal. %d %d", a, b)
	}
}

func DeepEqual(t *testing.T, a, b interface{}) {

}
