package cpb

import "reflect"

const (
	TypeInvalid uint64 = iota
	TypeBool
	TypeInt
	TypeString
	TypeStruct
	TypeJSON
	TypeMapString
	TypeBytes
	TypeFixed32
	TypeFixed64
	TypeSliceStruct
	TypeSliceInt
	TypeSliceString
	TypeSliceBytes
)

var marshalerType = reflect.TypeOf((*Marshaler)(nil)).Elem()
