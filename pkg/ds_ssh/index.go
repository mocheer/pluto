package ds_ssh

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

// SSHClient
type SSHClient struct {
	*ssh.Client
}

func New(user, password, host string) (*SSHClient, error) {
	// SSH 连接配置
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// SSH 连接
	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to remote server: %v", err)
	}
	return &SSHClient{Client: client}, nil
}

// Close 每次操作完都要关闭
func (m *SSHClient) Close() {
	m.Client.Close()
}

func (m *SSHClient) ReadFile(filePath string) ([]byte, error) {
	// 打开会话，注意，这个会话只能执行一次命令
	session, err := m.Client.NewSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()
	// 读取文件内容
	output, err := session.Output("cat " + filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read remote file: %v", err)
	}
	return output, nil
}

// GetFilenames 获取指定目录及所有子目录下的所有文件
func (m *SSHClient) GetFilenames(dir string) (files []string, err error) {
	// 打开会话，注意，这个会话只能执行一次命令
	session, err := m.Client.NewSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()
	// 执行 ls 命令获取文件列表
	command := fmt.Sprintf("ls -p %s | grep -v /", strings.TrimRight(dir, "/"))
	output, err := session.Output(command)
	if err != nil {
		return nil, fmt.Errorf("failed to read remote files: %v", err)
	}

	for _, file := range strings.Split(string(output), "\n") {
		if file != "" {
			files = append(files, file)
		}
	}

	return files, nil
}

// Remove
func (m *SSHClient) Remove(name string) (err error) {
	// 打开会话，注意，这个会话只能执行一次命令
	session, err := m.Client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()
	command := fmt.Sprintf("rm %s", name)
	err = session.Run(command)
	if err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}
	return nil
}
