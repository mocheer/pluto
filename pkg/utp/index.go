package utp

import (
	"log"
	"net/url"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
	"github.com/mocheer/pluto/pkg/ds"
)

// Save 请求http数据接口
func Save(uri string, fileName string, orgin string) error {
	return Request(uri, orgin, func(resp *colly.Response) {
		// 确保目录存在
		ds.CreateDir(fileName)
		if err := resp.Save(fileName); err != nil {
			log.Fatal(err)
		}
	})
}

// Request 请求http数据接口
func Request(uri string, orgin string, callback func(resp *colly.Response)) error {
	c := colly.NewCollector(
		// colly.DetectCharset(),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.134 Safari/537.36 Edg/103.0.1264.71"),
	)

	if p, err := proxy.RoundRobinProxySwitcher(
		"218.85.200.24:9999",
		"112.25.12.39:80",
		"118.114.250.104:8888",
		"220.176.168.17:8888",
		"125.86.191.189:8888",
		"180.103.191.96:8888",
		"120.41.249.122:8080",
		"220.173.123.212:9999",
		"119.5.184.239:8888",
		"110.83.12.171:8888",
		"182.34.222.222:8888",
		"58.52.82.141:8080",
		"101.16.64.235:8888",
		"119.52.50.27:8888",
		"123.171.1.224:9999",
		"113.13.177.25:9999",
		"112.195.240.214:8080",
	); err == nil {
		c.SetProxyFunc(p)
	}
	// extensions.RandomUserAgent(c)
	// extensions.Referer(c)

	if orgin == "" {
		orgin = uri
	}
	u, err := url.Parse(orgin)
	if err != nil {
		return err
	}

	c.OnRequest(func(r *colly.Request) {
		// host := u.Host
		host := "localhost:5173"
		r.Headers.Set("Host", host)
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*") //
		r.Headers.Set("Origin", host)
		r.Headers.Set("Referer", u.String())
		//关键头 如果没有 则返回 错误
		r.Headers.Set("Accept-Encoding", "gzip,deflate")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
	})
	c.OnResponse(callback)
	// c.OnError(func(r *colly.Response, err error) {
	// 	fmt.Println(err)
	// })
	err = c.Visit(uri)
	c.Wait() //等待结束
	return err
}

// GetBody
func GetBody(uri string, orgin string) ([]byte, error) {
	var data []byte
	err := Request(uri, orgin, func(resp *colly.Response) {
		data = resp.Body
	})
	return data, err
}
