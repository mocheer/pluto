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

func TestProxy(t *testing.T) {
	// @see https://www.zdaye.com/dayProxy.html
	urls := []string{
		"http://1.15.156.141:7890",
		"http://113.57.84.39:9091",
		"http://111.225.152.219:8089",
		"http://116.205.229.85:80",
		"http://113.223.214.53:8089",
		"http://1.13.15.155:3128",
		"http://101.230.172.86:9443",
		"http://8.219.97.248:80",
		"http://120.206.182.34:3128",
		"http://175.8.132.228:7890",
	}
	for _, url := range urls {
		c := ctp.New()
		c.SetRequestTimeout(time.Second * 3)
		c.SetProxies(url)
		data, err := c.Get("https://qifu-api.baidubce.com/ip/local/geo/v1/district") // https://www.baidu.com
		t.Log(url, string(data), err)
	}

}

func TestProxy2(t *testing.T) {
	ips := []string{
		"http://117.71.149.218:8089",
		"http://114.231.45.160:8888",
		"http://117.71.149.151:8089",
		"http://43.138.208.113:8080",
		"http://222.190.208.37:8089",
		"http://113.223.215.50:8089",
		"http://111.225.152.187:8089",
		"http://117.69.237.21:8089",
		"http://117.71.149.162:8089",
		"http://117.57.93.185:8089",
		"http://117.94.126.129:9000",
		"http://182.34.22.62:8089",
		"http://175.8.132.228:7890",
		"http://183.165.244.5:8089",
		"http://49.85.15.27:9000",
	}
	sips := []string{}
	for _, ip := range ips {
		c := ctp.New()
		c.SetRequestTimeout(time.Second * 1)
		c.SetProxies(ip)
		data, err := c.Get("https://qifu-api.baidubce.com/ip/local/geo/v1/district") // https://www.baidu.com
		if err == nil {
			t.Log(string(data))
			sips = append(sips, ip)
		} else {
			t.Log(ip, err)
		}
	}
	fmt.Printf("%+v", sips)
}
