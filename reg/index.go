package reg

// 正则表达式
var (
	// Brace 匹配花括号内的字符串`{xxx}`，常用于字符串格式化替换
	Brace = `{([^}]+)}`
	// CamelCase 匹配短连接字符`-`，常用于转成驼峰大写
	CamelCase = `-+(.)?`
	// CommandParams 匹配命令行参数，用于按空格分割（不包括引号内的空格）
	CommandParams = `"([^"]*?)"|(\S+)`
)
