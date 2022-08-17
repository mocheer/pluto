package sys

import (
	"bufio"
	"fmt"

	"github.com/amenzhinsky/go-memexec"
)

// MemExec
// 在内存中直接执行程序，不需要在硬盘临时创建文件
func MemExec(bs []byte, params ...string) error {
	exe, err := memexec.New(bs)
	if err != nil {
		return err
	}
	defer exe.Close()

	command := exe.Command(params...)

	stdout, err := command.StdoutPipe()
	command.Stderr = command.Stdout

	if err != nil {
		return err
	}
	err = command.Start()
	if err != nil {
		return err
	}
	// 从管道中实时获取输出并打印到终端
	scanner := bufio.NewScanner(stdout)
	// scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	//
	if err = command.Wait(); err != nil {
		return err
	}
	return nil
}
