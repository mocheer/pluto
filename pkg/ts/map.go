package ts

import (
	"encoding/json"

	"github.com/mocheer/pluto/pkg/fn"
)

// Map
type Map[T any] map[string]T

// 获取map的key值集合
func (m Map[T]) Keys() []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// 获取map的key值个数
func (m Map[T]) Len() int {
	return len(m)
}

//
func (m Map[T]) Bytes() []byte {
	bytes, _ := json.Marshal(m)
	return bytes
}

//
func (m Map[T]) String() string {
	return fn.B2S(m.Bytes())
}

// 合并两个map
func (m Map[T]) Assign(m2 Map[T]) Map[T] {
	for k, v := range m2 {
		m[k] = v
	}
	return m
}
