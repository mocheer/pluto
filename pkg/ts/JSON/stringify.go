package JSON

import (
	"encoding/json"

	"github.com/mocheer/pluto/pkg/fn"
)

// Stringify json反序列化
func Stringify(v any) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return fn.BytesToString(bytes)
}
