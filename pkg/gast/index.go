package gast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

type FuncCode string

func (m FuncCode) TrimFunc() FuncCode {
	return FuncCode(strings.TrimPrefix(string(m), "func "))
}

func (m FuncCode) String() string {
	return string(m)
}

// ExtractFunctionCode 从指定的Go文件中提取指定函数的代码内容
func ExtractFunctionCode(filePath string, functionName string) (FuncCode, error) {
	// 打开文件
	fileSet := token.NewFileSet()
	fileAST, err := parser.ParseFile(fileSet, filePath, nil, parser.ParseComments)
	if err != nil {
		return "", fmt.Errorf("failed to parse file: %v", err)
	}

	// 遍历AST，查找指定函数
	for _, decl := range fileAST.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		if funcDecl.Name.Name == functionName {
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
	}

	return "", fmt.Errorf("function %s not found in file", functionName)
}
