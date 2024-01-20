package object_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/mocheer/pluto/pkg/ts/object"
	"github.com/stretchr/testify/assert"
	"golang.org/x/image/tiff"
)

func TestGetKind(t *testing.T) {
	assert.Equal(t, object.GetKind(1), reflect.Int)
	assert.Equal(t, object.GetKind(1.0), reflect.Float64)
	assert.Equal(t, object.GetKind(nil), reflect.Invalid)
	assert.Equal(t, object.GetKind("string"), reflect.String)
	assert.Equal(t, object.GetKind(map[string]string{}), reflect.Map)
	assert.Equal(t, object.GetKind(map[string]int{}), reflect.Map)
	assert.Equal(t, object.GetKind(false), reflect.Bool)
}
func TestIsTime(t *testing.T) {
	t.Log(object.GetType(time.Now())) //struct
}

func TestType(t *testing.T) {
	type User struct {
		ID string
	}
	u := &User{ID: ""}
	t.Log(reflect.ValueOf(u).Elem().Type() == reflect.TypeOf(u).Elem())
}

func TestGetFuncName(t *testing.T) {
	t.Log(object.GetFuncName(reflect.ValueOf))
}

func TestGetTypeName(t *testing.T) {
	t.Log(object.GetTypeName(time.Timer{}))
	t.Log(object.GetTypeName(&time.Timer{}))

	t.Log(object.GetTypeName(time.Layout))
	// t.Log(object.GetTypeName(&time.Layout))

	//
	t.Log(object.GetPackageName(time.Timer{}))
	t.Log(object.GetPackageName(tiff.CCITTGroup3))
	// t.Log(object.GetPackageName(&time.Layout))
}
