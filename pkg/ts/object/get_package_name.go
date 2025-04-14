package object

import (
	"reflect"
	"strings"
)

// GetPackageName 获取包名
// 无法获取到预声明的string和error，以及匿名变量的包名
// 能够获取到各种类型的包名，如结构体、整型、浮点型等
// 这里的包名是完整的路径
func GetPackageName(v any) string {
	return reflect.TypeOf(v).PkgPath()
}

func GetPackageIdentName(v any) string {
	s := GetPackageName(v)
	return s[strings.LastIndex(s, "/")+1:]
}
