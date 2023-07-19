package object_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/mocheer/pluto/pkg/ts/object"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
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

func TestKeys(t *testing.T) {
	dj := gjson.Parse(`{"data":null,"int":0,"bool":false,"key":"key"}`).Map()
	fields := object.Keys(dj)
	t.Log(fields)
}
