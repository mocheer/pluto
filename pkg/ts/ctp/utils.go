package ctp

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

// CheckProxy 用于测试代理地址是否可访问
func CheckProxy(proxyAddr string, targetAddr string) error {
	c := colly.NewCollector()

	if p, err := proxy.RoundRobinProxySwitcher(proxyAddr); err == nil {
		c.SetProxyFunc(p)
	}

	// 访问目标地址，不能访问时会抛出错误，当然，这种方式不准确，因为可能是目标服务器出现了问题
	return c.Visit(targetAddr)
}

func Save(uri string, fileName string) error {
	return New().Save(uri, fileName)
}

func Get(uri string) ([]byte, error) {
	return New().Get(uri)
}
