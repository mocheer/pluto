// +build windows

package cmd

import (
	"os"
	"path/filepath"
)

const (
	dllFileProtocolHandler = "url.dll,FileProtocolHandler"
	dllSHExitWindowsEx     = "shell32.dll,SHExitWindowsEx"
	dllControlRunDLL       = "shell32.dll,Control_RunDLL"
	dllLockWorkStation     = "user32.dll,LockWorkStation"
)

var (
	runDll32 = filepath.Join(os.Getenv("SYSTEMROOT"), "System32", "rundll32.exe")
)

//Open 打开网页、文件、目录,相对路径C:\WINDOWS\System32\
//不能打开环境变量指向的目录文件？
func Open(params []string) error {
	params = append([]string{dllFileProtocolHandler}, params...)
	return Exec(runDll32, params...)
}

// Coldboot 冷启动 windows Explorer
func Coldboot() error {
	return Exec(runDll32, dllSHExitWindowsEx, "-1")
}

//Logoff 注销
func Logoff() error {
	return Exec(runDll32, dllSHExitWindowsEx, "0")
}

//Shutdown 关机
func Shutdown() error {
	return Exec(runDll32, dllSHExitWindowsEx, "1")
}

//Reboot 重启
func Reboot() error {
	return Exec(runDll32, dllSHExitWindowsEx, "2")
}

//Poweroff 关闭电源
func Poweroff() error {
	return Exec(runDll32, dllSHExitWindowsEx, "4")
}

//LockWork 锁定计算机
func LockWork() error {
	return Exec(runDll32, dllLockWorkStation)
}

//ControlRunDLL 控制面板
func ControlRunDLL() error {
	return Exec(runDll32, dllControlRunDLL)
}

//Start 启动程序，包括url，支持相对路径,但会显示一个控制台
func Start(params []string) error {
	params = append([]string{"/C", "start"}, params...)
	return Exec("cmd", params...)
}

// RegisterService TODO
func RegisterService(name string, path string, start string) error {
	return Exec("cmd", "/C", "sc", "create", name, `binPath= ""`+path+`""`, "start= "+start)
}
