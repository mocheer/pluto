package gast

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/mocheer/pluto/pkg/ds"
)

type FuncCode string

func (m FuncCode) TrimFunc() FuncCode {
	return FuncCode(strings.TrimPrefix(string(m), "func "))
}

func (m FuncCode) String() string {
	return string(m)
}

// ExtractFuncCode 从指定的Go文件中提取指定函数的代码内容
func ExtractFuncCode(filePath string, functionName string) (FuncCode, error) {
	fileSet, funcDecl, err := ExtractFunc(filePath, functionName)
	if err != nil {
		return "", err
	}

	// 找到函数，提取代码范围
	start := fileSet.Position(funcDecl.Pos()).Offset
	end := fileSet.Position(funcDecl.End()).Offset
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file content: %v", err)
	}
	code := string(fileContent[start:end])
	// 去掉开头的 "func" 关键字
	return FuncCode(code), nil

}

// ExtractFuncCode 从指定的Go文件中提取指定函数的代码内容
func ExtractFunc(filePath string, funName string) (*token.FileSet, *ast.FuncDecl, error) {
	// 打开文件
	fileSet := token.NewFileSet()
	fileAST, err := parser.ParseFile(fileSet, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse file: %v", err)
	}

	// 遍历AST，查找指定函数
	for _, decl := range fileAST.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		if funcDecl.Name.Name == funName {
			return fileSet, funcDecl, nil
		}
	}

	return fileSet, nil, fmt.Errorf("function %s not found in file", funName)
}

// ExtractImportPackages 获取一个go文件的内部所有引用的第三方包
func ExtractImportPackages(filePath string) ([]*ast.ImportSpec, error) {
	// 打开文件
	fileSet := token.NewFileSet()
	fileAST, err := parser.ParseFile(fileSet, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file: %v", err)
	}
	return fileAST.Imports, nil
}

// GetModelMethod
// 这里只支持结构体和函数，后续应该考虑支持自定义类型，很多自定义类型会有自己封装的方法，即使原始类型是int、float64、string也有一堆自己的方法
func GetModelMethod(v interface{}) (method *DIYMethods, err error) {
	method = new(DIYMethods)

	//
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Func:
		fullPath := runtime.FuncForPC(value.Pointer()).Name()
		err = method.parserPath(fullPath)
		if err != nil {
			return nil, err
		}
	case reflect.Struct:
		method.pkgPath = value.Type().PkgPath()
		method.BaseStructType = value.Type().Name()
	default:
		// fmt.Println(value.Kind())
		return nil, fmt.Errorf("method param must be a function or struct")
	}

	var p *build.Package

	// if struct in main file
	ctx := build.Default
	if method.pkgPath == "main" {
		var skip int
		var file string
		for {
			_, file, _, _ = runtime.Caller(skip)
			if !(strings.Contains(file, "gorm/gen/generator.go") || strings.Contains(file, "gorm/gen/internal")) || file == "" {
				break
			}
			skip++
		}
		p, err = ctx.ImportDir(filepath.Dir(file), build.ImportComment)
	} else {
		p, err = ctx.Import(method.pkgPath, "", build.ImportComment)
	}
	if err != nil {
		return nil, fmt.Errorf("diy method dir not found:%s.%s %w", method.pkgPath, method.MethodName, err)
	}

	for _, file := range p.GoFiles {
		goFile := p.Dir + "/" + file
		if ds.IsExist(goFile) {
			method.pkgFiles = append(method.pkgFiles, goFile)
		}
	}
	if len(method.pkgFiles) == 0 {
		return nil, fmt.Errorf("diy method file not found:%s.%s", method.pkgPath, method.MethodName)
	}

	// read files got methods
	return method, method.LoadMethods()
}
