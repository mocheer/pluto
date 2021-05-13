package JSON

import (
	"encoding/json"

	"github.com/mocheer/pluto/fn"
)

// Parse json 解析 序列化
func Parse(data string) map[string]interface{} {
	var v map[string]interface{}
	err := json.Unmarshal(fn.StringBytes(data), &v)
	if err != nil {
		panic(err)
	}
	return v
}
