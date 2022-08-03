package web_request

import (
	"log"
	"net/url"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/mocheer/pluto/pkg/ds"
)

// Save 请求http数据接口
func Save(uri string, fileName string, orgin string) error {
	c := colly.NewCollector(
	// colly.DetectCharset(),
	// colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.134 Safari/537.36 Edg/103.0.1264.71"),
	)

	extensions.RandomUserAgent(c)
	// extensions.Referer(c)

	if orgin == "" {
		orgin = uri
	}
	u, err := url.Parse(orgin)
	if err != nil {
		return err
	}

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Host", u.Host)
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*") //
		r.Headers.Set("Origin", u.Host)
		r.Headers.Set("Referer", u.String())
		//关键头 如果没有 则返回 错误
		r.Headers.Set("Accept-Encoding", "gzip,deflate")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
	})
	c.OnResponse(func(resp *colly.Response) {
		// 确保目录存在
		ds.CreateDir(fileName)
		if err := resp.Save(fileName); err != nil {
			log.Fatal(err)
		}
	})
	// c.OnError(func(r *colly.Response, err error) {
	// 	fmt.Println(err)
	// })

	err = c.Visit(uri)
	c.Wait() //等待结束
	return err
}
