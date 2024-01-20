package object

import (
	"reflect"
	"runtime"
)

// GetFuncName
// 包括包名，这里的包名是引用时的别名，一般是最后一个目录名称
func GetFuncName(fn any) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}
