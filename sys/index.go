package sys

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// GetCurrentPath 获取当前程序所在位置
// go install 安装后在任意位置执行时，GetCurrentPath()= `%go%/bin`
func GetCurrentPath() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

// 获取当前函数所在文件路径（一旦文件编译生成）
// skip=0 是当前sys的目录
// skip=1 是当前调用 GetCurrentFuncPath(1) 的目录
func GetCurrentFuncPath(skip int) string {
	// runtime.Caller 用于打印调用栈信息
	_, filename, _, ok := runtime.Caller(skip)
	if ok {
		return path.Dir(filename)
	} else {
		panic(`error:GetCurrentFuncPath`)
	}

}
