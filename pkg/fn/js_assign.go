package fn

import (
	"encoding/json"
	"fmt"
)

// Assign
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
