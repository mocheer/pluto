package ctp_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/ts/ctp"
)

func TestGet(t *testing.T) {
	// 16554181e1d8f9f3b82ce84fe953c164
	// 0b79a07d2808103ab84aa56485c331a8
	err := ctp.Save(`https://t5.tianditu.gov.cn/vec_w/wmts?SERVICE=WMTS&REQUEST=GetTile&VERSION=1.0.0&LAYER=vec&STYLE=default&TILEMATRIXSET=w&FORMAT=tiles&TILECOL=4&TILEROW=3&TILEMATRIX=3&tk=0b79a07d2808103ab84aa56485c331a8`, "test.jpg")

	t.Log(err)

}

func TestProxy(t *testing.T) {
	urls := []string{
		"36.111.146.161:9000",
		"36.6.145.98:8089",
		"34.168.233.208:8585",
		"43.156.241.242:8089",
		"146.59.2.185:80",
	}
	for _, url := range urls {
		t.Log(url, ctp.CheckProxy(url, "https://www.baidu.com"))
	}

}
