package cmd

import (
	"os/exec"
)

// LookPath 根据所给的二进制文件名file在PATH环境变量中寻找此二进制文件所在的路径信息
// Command  返回用于执行name代表的命令的Cmd结构，Cmd中设置了Path和Args，其它fields都为空。
// CombinedOutput 运行命令c，并返回combined output和error。
// OutPut 运行指定的命令，并返回标准输出
// Run 运行指定的命令，并等待其完成。
// Start 运行一个指定的命令c，但不等待其完成。
// Wait 等待命令c退出，命令c必须是由Start方法运行的。

// Exec 执行命令并返回结果。命令行参数在窗口输入的时候需要带引号，但这里的参数不需要，反而要去掉
func Exec(name string, params ...string) string {
	command := exec.Command(name, params...)
	result, err := command.Output()
	if err != nil {
		return err.Error()
	}
	return string(result)
}
