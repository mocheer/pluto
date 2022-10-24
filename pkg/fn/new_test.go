package fn_test

import (
	"fmt"
	"testing"

	"github.com/mocheer/pluto/pkg/fn"
)

func TestNewSlices(t *testing.T) {
	type User struct {
		Name string `json:"name"`
	}
	u := User{}
	users := fn.NewSlice(u)
	fmt.Println(users)
}
