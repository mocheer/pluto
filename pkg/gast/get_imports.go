package gast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"slices"
	"strings"
)

// 获取文件导入的包
func GetImports(filePath string) []string {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		panic(fmt.Errorf("解析文件[%s]错误:%s", filePath, err))
	}
	return getImports(file)
}

// 获取文件导入的包
func getImports(file *ast.File) []string {
	imports := make([]string, len(file.Imports))
	for i, imp := range file.Imports {
		imports[i] = strings.Trim(imp.Path.Value, `"`)
	}
	return imports
}

// 分析函数引用的第三方包
func GetFunctionImports(filePath string, targetFuncName string) []string {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		panic(fmt.Errorf("解析文件[%s]错误:%s", filePath, err))
	}
	// 遍历文件中的所有声明，找到目标函数
	var targetFunc *ast.FuncDecl
	ast.Inspect(file, func(n ast.Node) bool {
		if funcDecl, ok := n.(*ast.FuncDecl); ok {
			if funcDecl.Name.Name == targetFuncName {
				targetFunc = funcDecl
				return false // 找到目标函数后停止遍历
			}
		}
		return true
	})

	if targetFunc == nil {
		fmt.Printf("未找到函数：%s\n", targetFuncName)
		return nil
	}
	return getFunctionImports(targetFunc, getImports(file))
}

// 分析函数引用的第三方包
// 这里需要拿到allImports，然后根据代码中的标识符，判断名称
func getFunctionImports(funcDecl *ast.FuncDecl, allImports []string) []string {
	usedPackages := []string{}
	allPackages := []string{}
	for _, p := range allImports {
		name := p
		if strings.Contains(name, "/") {
			name = name[(strings.LastIndex(name, "/") + 1):]
		}
		allPackages = append(allPackages, name)
	}
	// 遍历函数体中的所有节点
	ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
		switch x := n.(type) {
		//
		case *ast.SelectorExpr:
			if ident, ok := x.X.(*ast.Ident); ok {

				if index := slices.Index(allPackages, ident.Name); index != -1 {
					name := allImports[index]
					if !slices.Contains(usedPackages, name) {
						usedPackages = append(usedPackages, name)
					}
				}

			}
		// 函数调用
		case *ast.CallExpr:
			if ident, ok := x.Fun.(*ast.Ident); ok {
				if index := slices.Index(allPackages, ident.Name); index != -1 {
					name := allImports[index]
					if !slices.Contains(usedPackages, name) {
						usedPackages = append(usedPackages, name)
					}
				}
			}
		}
		return true
	})

	return usedPackages
}
