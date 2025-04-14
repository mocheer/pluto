package gast

import (
	"fmt"
	"go/ast"
	"log"
	"strings"
)

// Param parameters in method
type Param struct { // (user model.User)
	PkgPath   string // package's path: internal/model
	Package   string // package's name: model
	Name      string // param's name: user
	Type      string // param's type: User
	IsArray   bool   // is array or not
	IsPointer bool   // is pointer or not
}

// Eq if param equal to another
func (p *Param) Eq(q Param) bool {
	return p.Package == q.Package && p.Type == q.Type
}

// IsError ...
func (p *Param) IsError() bool {
	return p.Type == "error"
}

// IsGenM ...
func (p *Param) IsGenM() bool {
	return p.Package == "gen" && p.Type == "M"
}

// IsGenRowsAffected ...
func (p *Param) IsGenRowsAffected() bool {
	return p.Package == "gen" && p.Type == "RowsAffected"
}

// IsMap ...
func (p *Param) IsMap() bool {
	return strings.HasPrefix(p.Type, "map[")
}

// IsGenT ...
func (p *Param) IsGenT() bool {
	return p.Package == "gen" && p.Type == "T"
}

// IsInterface ...
func (p *Param) IsInterface() bool {
	return p.Type == "interface{}"
}

// IsNull ...
func (p *Param) IsNull() bool {
	return p.Package == "" && p.Type == "" && p.Name == ""
}

// InMainPkg ...
func (p *Param) InMainPkg() bool {
	return p.Package == "main"
}

// IsTime ...
func (p *Param) IsTime() bool {
	return p.Package == "time" && p.Type == "Time"
}

// IsSQLResult ...
func (p *Param) IsSQLResult() bool {
	return (p.Package == "sql" && p.Type == "Result") || (p.Package == "gen" && p.Type == "SQLResult")
}

// IsSQLRow ...
func (p *Param) IsSQLRow() bool {
	return (p.Package == "sql" && p.Type == "Row") || (p.Package == "gen" && p.Type == "SQLRow")
}

// IsSQLRows ...
func (p *Param) IsSQLRows() bool {
	return (p.Package == "sql" && p.Type == "Rows") || (p.Package == "gen" && p.Type == "SQLRows")
}

// SetName ...
func (p *Param) SetName(name string) {
	p.Name = name
}

// TypeName ...
func (p *Param) TypeName() string {
	if p.IsArray {
		return "[]" + p.Type
	}
	return p.Type
}

// TmplString param to string in tmpl
func (p *Param) TmplString() string {
	var res strings.Builder
	if p.Name != "" {
		res.WriteString(p.Name)
		res.WriteString(" ")
	}

	if p.IsArray {
		res.WriteString("[]")
	}
	if p.IsPointer {
		res.WriteString("*")
	}
	if p.Package != "" {
		res.WriteString(p.Package)
		res.WriteString(".")
	}
	res.WriteString(p.Type)
	return res.String()
}

// IsBaseType judge whether the param type is basic type
func (p *Param) IsBaseType() bool {
	switch p.Type {
	case "string", "byte":
		return true
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return true
	case "float64", "float32":
		return true
	case "bool":
		return true
	case "time.Time":
		return true
	default:
		return false
	}
}

func (p *Param) astGetParamType(param *ast.Field) {
	switch v := param.Type.(type) {
	case *ast.Ident:
		p.Type = v.Name
		if v.Obj != nil {
			p.Package = "UNDEFINED" // set a placeholder
		}
	case *ast.SelectorExpr:
		p.astGetEltType(v)
	case *ast.ArrayType:
		p.astGetEltType(v.Elt)
		p.IsArray = true
	case *ast.Ellipsis:
		p.astGetEltType(v.Elt)
		p.IsArray = true
	case *ast.MapType:
		p.astGetMapType(v)
	case *ast.InterfaceType:
		p.Type = "interface{}"
	case *ast.StarExpr:
		p.IsPointer = true
		p.astGetEltType(v.X)
	default:
		log.Fatalf("unknow param type: %+v", v)
	}
}

func (p *Param) astGetEltType(expr ast.Expr) {
	switch v := expr.(type) {
	case *ast.Ident:
		p.Type = v.Name
		if v.Obj != nil {
			p.Package = "UNDEFINED"
		}
	case *ast.SelectorExpr:
		p.Type = v.Sel.Name
		p.astGetPackageName(v.X)
	case *ast.MapType:
		p.astGetMapType(v)
	case *ast.StarExpr:
		p.IsPointer = true
		p.astGetEltType(v.X)
	case *ast.InterfaceType:
		p.Type = "interface{}"
	case *ast.ArrayType:
		p.astGetEltType(v.Elt)
		p.Type = "[]" + p.Type
	default:
		log.Fatalf("unknow param type: %+v", v)
	}
}

func (p *Param) astGetPackageName(expr ast.Expr) {
	switch v := expr.(type) {
	case *ast.Ident:
		p.Package = v.Name
	}
}

func (p *Param) astGetMapType(expr *ast.MapType) {
	p.Type = fmt.Sprintf("map[%s]%s", astGetType(expr.Key), astGetType(expr.Value))
}

func astGetType(expr ast.Expr) string {
	switch v := expr.(type) {
	case *ast.Ident:
		return v.Name
	case *ast.InterfaceType:
		return "interface{}"
	}
	return ""
}
