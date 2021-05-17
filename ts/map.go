package ts

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
