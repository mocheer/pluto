package ctp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/mocheer/pluto/pkg/ts/ctp"
)

func TestGet(t *testing.T) {
	// 16554181e1d8f9f3b82ce84fe953c164
	// 0b79a07d2808103ab84aa56485c331a8
	err := ctp.Save(`https://t5.tianditu.gov.cn/vec_w/wmts?SERVICE=WMTS&REQUEST=GetTile&VERSION=1.0.0&LAYER=vec&STYLE=default&TILEMATRIXSET=w&FORMAT=tiles&TILECOL=4&TILEROW=3&TILEMATRIX=3&tk=0b79a07d2808103ab84aa56485c331a8`, "test.jpg")

	t.Log(err)

}

func TestProxy2(t *testing.T) {
	ips := []string{
		"http://123.126.158.50:80",
		"http://101.34.72.57:7890",
		"http://180.103.127.9:7890",
		"http://36.138.56.214:3129",
		"http://36.139.164.147:8888",
		"http://139.198.168.65:7890",
		"http://62.234.182.56:443",
		"http://123.126.158.50:80",
		"http://47.99.180.88:7890",
		"http://139.196.78.175:7890",
		"http://111.177.63.86:8888",
		"http://47.106.144.184:7890",
		"http://81.68.190.184:3128",
	}
	sips := []string{}
	for _, ip := range ips {
		c := ctp.New()
		c.SetRequestTimeout(time.Second * 1)
		c.SetProxies(ip)
		data, err := c.Get("https://qifu-api.baidubce.com/ip/local/geo/v1/district") // https://www.baidu.com
		// data, err = c.Get("https://qifu-api.baidubce.com/ip/local/geo/v1/district")  // https://www.baidu.com
		if err == nil {
			t.Log(string(data))
			sips = append(sips, ip)
		} else {
			t.Log(ip, err)
		}
	}
	fmt.Printf("%+v", sips)
}
