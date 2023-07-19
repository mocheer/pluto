package ctp

func Save(uri string, fileName string) error {
	return New().Save(uri, fileName)
}

func Get(uri string) ([]byte, error) {
	return New().Get(uri)
}

// GetByProxy 用于测试代理地址是否可访问，不能访问时会抛出错误，当然，这种方式不准确，因为可能是目标服务器出现了问题
func GetByProxy(proxyAddr string, targetAddr string) error {
	return New().SetProxies(proxyAddr).Visit(targetAddr)
}
