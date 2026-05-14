package object

import (
	"encoding/json"
	"fmt"
)

// Assign
// 用json的序列化和反序列化来实现对象的拷贝
// 性能较差，不建议在性能敏感场景下使用
func Assign(a, b any) any {
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
