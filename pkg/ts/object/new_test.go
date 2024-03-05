package object_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/mocheer/pluto/pkg/ts/object"
)

func TestNew(t *testing.T) {
	type User struct {
		Name string `json:"name"`
	}
	u := User{Name: "Mocheer"}
	t.Log(object.GetKind(u), object.GetKind(object.New(u)))
	t.Log(object.GetKind(1), object.GetKind(object.New(1)))
	a, ok := object.New(u).(*User)
	a.Name = "mocheer"
	t.Log(*a, ok)
	t.Log(*a == u)
}

func TestNewSlices(t *testing.T) {
	type User struct {
		Name string `json:"name"`
	}
	u := User{Name: "Mocheer"}
	users := object.NewSlice(u)
	fmt.Println(users)
	//
	v := reflect.ValueOf(users).Elem()
	fmt.Println(v.Len())
	for i := 0; i < v.Len(); i++ {
		e := v.Index(i).Interface()
		t.Log(e)
	}
}
