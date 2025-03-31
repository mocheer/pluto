package gast

import (
	"testing"
)

func TestF(t *testing.T) {
	code, err := ExtractFunctionCode("./a_test.go", "TestF")
	if err != nil {
		t.Log(err)
	}
	t.Log(code)
	// t.Log(strings.TrimSpace(code))
}
