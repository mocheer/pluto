package fetch_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/mocheer/pluto/pkg/ts/fetch"
)

func TestFetch(t *testing.T) {
	resp, err := fetch.Fetch("https://baidu.com")
	if err != nil {
		panic(err)
	}
	fmt.Println("Status:", resp.StatusCode)
	// fmt.Println("Body:", resp.Text()) // 字符串输出
}

func TestFetchOptions(t *testing.T) {
	resp, err := fetch.Fetch("https://baidu.com",
		// fetch.WithProxy("http://127.0.0.1:7890"),
		fetch.WithTimeout(5*time.Second),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Status:", resp.StatusCode)
	// fmt.Println("Body:", resp.Text()) // 字符串输出
}
