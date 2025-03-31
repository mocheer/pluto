package gen

import (
	. "github.com/dave/jennifer/jen"
	"github.com/samber/lo"
)

// StructSchema 结构体的结构描述
type StructSchema struct {
	Name   string
	Fields []StructField
}

// StructField
type StructField struct {
	Name     string
	TypeName string
	JsonTag  string
	Comment  string
}

func (m StructSchema) Type() *Statement {
	codes := lo.Map(m.Fields, func(item StructField, _ int) Code {
		return createStructFieldStatement(item.Name, item.TypeName)
	})
	return Type().Id(m.Name).Struct(codes...)
}

// New
//
//	&name{
//		Age:  1,
//		Name: "a",
//	}
func (m StructSchema) New(params Dict) *Statement {
	return Op("&").Id(m.Name).Values(params)
}

func (m StructSchema) AddMethod(fn any) *Statement {

	// gast.ExtractFunctionCode
	return nil
}

func createStructFieldStatement(name string, typeName string) *Statement {
	s := Id(name)
	switch typeName {
	case "int":
		s.Int()
	case "string":
		s.String()
	case "uint8":
		s.Uint8()
	}
	return s
}
