package gen

import . "github.com/dave/jennifer/jen"

// NewMap 未完成
func NewMap(typ Code, params Dict) *Statement {
	return Map(typ).Any()
}

// NewMapStringAny
func NewMapStringAny(name string, params Dict) *Statement {
	return Map(String()).Any()
}
