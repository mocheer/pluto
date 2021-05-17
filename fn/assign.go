package fn

import (
	"encoding/json"
	"fmt"

	"github.com/mocheer/pluto/ts"
)

// Assign
func Assign(a, b interface{}) interface{} {
	jb, err := json.Marshal(b)
	if err != nil {
		fmt.Println("Marshal error b:", err)
	}
	err = json.Unmarshal(jb, &a)
	if err != nil {
		fmt.Println("Unmarshal error b-a:", err)
	}

	return a
}

// 合并两个map
func AssignMap(m1 ts.Map, m2 ts.Map) ts.Map {
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}
