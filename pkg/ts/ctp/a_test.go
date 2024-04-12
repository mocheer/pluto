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
		"http://117.160.250.132:80", "http://223.100.178.167:9091", "http://101.231.64.89:8443", "http://114.231.42.85:8089", "http://111.59.4.88:9002", "http://124.222.9.200:443", "http://221.6.139.190:9002", "http://183.234.215.11:8443", "http://119.3.14.229:8080", "http://47.106.144.184:7890", "http://111.26.177.28:9091", "http://39.104.142.132:8080", "http://39.105.51.125:80", "http://47.96.145.14:8888", "http://47.93.121.200:80", "http://117.70.49.102:8089", "http://114.231.45.145:8089", "http://111.16.50.12:9002", "http://112.30.155.83:12792", "http://47.100.254.82:80",
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
