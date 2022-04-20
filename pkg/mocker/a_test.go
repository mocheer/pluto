package mocker_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/mocker"
)

// TestRandString 测试生成随机字符串
func TestRandString(t *testing.T) {
	result := mocker.String(6)
	t.Log(result)
}

func TestRandName(t *testing.T) {
	result := mocker.Name()
	t.Log(result)
}
