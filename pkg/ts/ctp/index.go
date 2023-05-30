package ctp

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/proxy"
	"github.com/mocheer/pluto/pkg/ds"
)

type Ctp struct {
	*colly.Collector
}

func New() *Ctp {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.35"),
	)
	return &Ctp{Collector: c}
}

// Get
func (m *Ctp) Get(uri string) ([]byte, error) {
	var data []byte
	err := m.Request(uri, "", func(resp *colly.Response) {
		data = resp.Body
	})
	return data, err
}

// Request 请求http数据接口
func (m *Ctp) Request(uri string, orgin string, callback func(resp *colly.Response)) error {
	// extensions.RandomUserAgent(c)
	extensions.Referer(m.Collector)

	if orgin == "" {
		orgin = uri
	}
	u, err := url.Parse(orgin)
	if err != nil {
		return err
	}

	m.OnRequest(func(r *colly.Request) {
		host := u.Host
		r.Headers.Set("Host", host)
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*") //
		r.Headers.Set("Origin", host)
		r.Headers.Set("Referer", u.String())
		//关键头 如果没有 则返回 错误
		r.Headers.Set("Accept-Encoding", "gzip,deflate")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
	})
	m.OnResponse(callback)
	// c.OnError(func(r *colly.Response, err error) {
	// 	fmt.Println(err)
	// })
	err = m.Visit(uri)
	m.Wait() //等待结束
	return err
}

// Save 请求http数据接口
func (m *Ctp) Save(uri string, fileName string) error {
	return m.Request(uri, "", func(resp *colly.Response) {
		// 确保目录存在
		ds.CreateDir(fileName)
		if err := resp.Save(fileName); err != nil {
			log.Fatal(err)
		}
	})
}

// EnableInsecureTLS 是否忽略证书校验企业微信的一些接口请求需要忽略证书校验
func (m *Ctp) EnableInsecureTLS(value bool) {
	// 创建一个忽略证书校验的transport
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: value},
	}
	m.WithTransport(tr)
}

// SetProxies
// SetProxies(
// "36.111.146.161:9000",
// "36.6.145.98:8089",
// "34.168.233.208:8585",
// "43.156.241.242:8089",
// "146.59.2.185:80",
// )
func (m *Ctp) SetProxies(addr ...string) {
	if p, err := proxy.RoundRobinProxySwitcher(addr...); err == nil {
		m.SetProxyFunc(p)
	}
}
