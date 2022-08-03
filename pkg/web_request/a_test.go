package web_request_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/web_request"
)

func TestGet(t *testing.T) {
	// 16554181e1d8f9f3b82ce84fe953c164
	// 0b79a07d2808103ab84aa56485c331a8
	err := web_request.Save(`https://t5.tianditu.gov.cn/vec_w/wmts?SERVICE=WMTS&REQUEST=GetTile&VERSION=1.0.0&LAYER=vec&STYLE=default&TILEMATRIXSET=w&FORMAT=tiles&TILECOL=4&TILEROW=3&TILEMATRIX=3&tk=0b79a07d2808103ab84aa56485c331a8`, "test.png", "https://www.tianditu.gov.cn/")

	t.Log(err)

}
