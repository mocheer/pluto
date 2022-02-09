package ts

import (
	"encoding/json"

	"github.com/mocheer/pluto/pkg/fn"
)

// Map map[string]interface{}的缩写
type Map map[string]interface{}

// 获取map的key值集合
func (m Map) Keys() []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// 获取map的key值个数
func (m Map) Len() int {
	return len(m)
}

//
func (m Map) Bytes() []byte {
	bytes, _ := json.Marshal(m)
	return bytes
}

//
func (m Map) String() string {
	return fn.B2S(m.Bytes())
}

// 合并两个map
func (m Map) Assign(m2 Map) Map {
	for k, v := range m2 {
		m[k] = v
	}
	return m
}
