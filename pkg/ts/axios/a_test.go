package axios_test

import (
	"net/http"
	"testing"

	"github.com/mocheer/pluto/pkg/ts/axios"
)

func TestGetAsync(t *testing.T) {
	// 测试 Get 方法
	for i := 0; i < 20; i++ {
		axios.Get("https://baidu.com").Then(func(resp *axios.Response) {
			if resp.StatusCode != http.StatusOK {
				t.Errorf("Get failed, status code: %d", resp.StatusCode)
			}
		}).Catch(func(err error) {
			t.Errorf("Get failed: %v", err)
		}).Finally(func() {
			t.Log(i)
		})
	}
}

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
