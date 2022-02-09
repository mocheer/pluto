package sys

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// GetCurrentPath 获取程序当前执行的路径 => 当前命令行执行的位置
func GetCurrentPath() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path
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

// GetExePath 获取程序所在位置
// go install 安装后在任意位置执行时，GetExePath()= `%go%/bin`
func GetExePath() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}
