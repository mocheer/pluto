package fetch

// New 创建一个 Fetch 请求客户端
// 功能一：不同于fetch直接请求，可以复用colly实例批量请求多个URL，提升性能
// 功能二：会自动管理并发请求数量，避免对目标服务器造成过大压力，避免本地网络资源拥堵
func New() *FetchClient {
	return &FetchClient{
		MaxConcurrency: 10,
	}
}
