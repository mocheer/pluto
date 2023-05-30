package object_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/mocheer/pluto/pkg/ts/object"
)

func TestNewSlices(t *testing.T) {
	type User struct {
		Name string `json:"name"`
	}
	u := User{Name: "Mocheer"}
	users := object.NewSlicePointer(u)
	fmt.Println(users)
	//
	v := reflect.ValueOf(users).Elem()
	fmt.Println(v.Len())
	for i := 0; i < v.Len(); i++ {
		e := v.Index(i).Interface()
		t.Log(e)
	}
}
