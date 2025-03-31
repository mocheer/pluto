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
		"http://39.108.2.73:3838", "http://47.100.67.65:7890", "http://120.25.1.15:7890",
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
