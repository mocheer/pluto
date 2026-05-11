package fetch

// Fetch 发起 HTTP 请求，返回标准化的响应或错误。
// 用法类似于 JS 的 fetch，支持通过 RequestOption 配置请求。
// TODO 目前colly即用即抛，后续考虑实现复用colly实例用于批量请求提升性能
// TODO 错误重试机制
func Fetch(u string, opts ...RequestOption) (*Response, error) {
	return New().Fetch(u, opts...)
}
