package reg

import (
	"regexp"
)

var (
	// CommandParams 命令行参数
	CommandParams = regexp.MustCompile(`"([^"]*?)"|(\S+)`)
)
