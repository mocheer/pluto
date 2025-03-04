package gen

import . "github.com/dave/jennifer/jen"

//
type StructSchema struct {
	Name   string
	Fields []StructField
}

//
type StructField struct {
	Name    string
	Type    string
	JsonTag string
	Comment string
}

func (m StructSchema) Define() *Statement {
	return Type().Id(m.Name).Struct()
}

// New
// &name{
// 	Age:  1,
// 	Name: "a",
// }
func (m StructSchema) New(params Dict) *Statement {
	return Op("&").Id(m.Name).Values(params)
}

// Struct
// type name struct{
//   params...
// }
func DefineStruct(name string, params ...Code) *Statement {
	return Type().Id(name).Struct(params...)
}
