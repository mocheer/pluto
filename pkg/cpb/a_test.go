package cpb_test

import (
	_ "embed"
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/mocheer/pluto/pkg/cpb"
	"github.com/vmihailenco/msgpack/v5"
	"google.golang.org/protobuf/encoding/protowire"
	"gorm.io/datatypes"
)

type User struct {
	Name        string
	Age         int
	D           int
	Anniversary map[string]string
	Data        []byte
	Num         int
	CreatedAt   *time.Time
}

var u = &User{
	Name: "mocheer",
	Age:  30,
	D:    -10,
	Anniversary: map[string]string{
		"birthday":  "19921214",
		"birthday2": "1214",
	},
	Data: []byte{1, 2, 3, 4, 5},
}

var us = []*User{u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u, u}

//go:embed marshal.go
var bs []byte

func Test1(t *testing.T) {

	data := cpb.Marshal(u)
	t.Log(data)
	t.Log(len(data))
	// t.Log(len(u.Anniversary))

	a := &User{}

	av := reflect.ValueOf(a)
	t.Log(av.Kind() == reflect.Pointer, av.IsNil(), av.IsZero(), av.IsValid(), av.Kind())
	avv := av.Elem()
	t.Log(avv.Kind() == reflect.Pointer, avv.IsZero(), avv.IsValid(), av.Type().Kind(), avv.Kind())
	a = nil
	av = reflect.ValueOf(a)
	t.Log(av.Kind() == reflect.Pointer, av.IsNil(), av.IsZero(), av.IsValid())
	cv := reflect.ValueOf(u).Elem().Field(4)
	t.Log(cv.Kind())
	t.Log(u.Num, cv.IsZero(), cv.IsValid())

	var b string
	bv := reflect.ValueOf(b)
	t.Log(bv.IsZero(), bv.IsValid())

}

func Test100(t *testing.T) {
	var a datatypes.JSON

	av := reflect.ValueOf(a)
	at := av.Type()
	var b []byte
	bv := reflect.ValueOf(b)

	t.Log(av.Kind(), bv.Kind(), reflect.Slice)
	t.Log(av.Kind() == reflect.Slice, bv.Kind() == reflect.Slice)

	_, ok1 := av.Interface().([]byte)
	t.Log(ok1)
	_, ok2 := bv.Interface().([]uint8)
	t.Log(ok2)

	t.Log(at.String())

	_, ok := av.Interface().(datatypes.JSON)
	t.Log(ok)

}

func Test200(t *testing.T) {
	var data1 []byte
	data1 = protowire.AppendVarint(data1, 0)
	data1 = protowire.AppendVarint(data1, 1)

	var data2 []byte
	data2 = protowire.AppendTag(data2, protowire.Number(3), 20)

	t.Log(data1)
	t.Log(data2)

	a, b, n := protowire.ConsumeTag(data2)
	t.Log(a, b, n)

}

func TestArray(t *testing.T) {
	var a = [6][]byte{
		{133},
	}
	t.Log(cpb.Marshal(a))

}

func BenchmarkCPB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		cpb.Marshal(us)
	}
}

func BenchmarkJSON(b *testing.B) {
	for n := 0; n < b.N; n++ {
		json.Marshal(us)
	}
}

func BenchmarkMessagePack(b *testing.B) {
	for n := 0; n < b.N; n++ {
		msgpack.Marshal(us)

	}

}

func BenchmarkCPB2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		cpb.Marshal(bs)

	}

}

func BenchmarkJSON2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		json.Marshal(bs)

	}
}

func BenchmarkMessagePack2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		msgpack.Marshal(bs)

	}

}
