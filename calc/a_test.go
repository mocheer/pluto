package calc_test

import (
	"testing"

	"github.com/mocheer/pluto/calc"
)

// TestRdString 测试随机字符串
func TestRdString(t *testing.T) {
	result := calc.String(6)
	t.Log(result)
}
