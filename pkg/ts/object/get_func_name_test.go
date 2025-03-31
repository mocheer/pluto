package object

import "testing"

func TestGetFuncName(t *testing.T) {
	t.Log(GetFuncName(TestGetFuncName))
}
