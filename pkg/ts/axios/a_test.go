package axios_test

import (
	"net/http"
	"testing"

	"github.com/mocheer/pluto/pkg/ts/axios"
)

// TestGet
func TestGetSync(t *testing.T) {
	// 测试 Get 方法
	resp, err := axios.GetSync("https://www.baidu.com")
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Get failed, status code: %d", resp.StatusCode)
	}
}
