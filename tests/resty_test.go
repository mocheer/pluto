package tests

import (
	"crypto/tls"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestResty1(t *testing.T) {

	resp, err := resty.New().R().Get("https://www.baidu.com/")
	assert.Equal(t, nil, err)
	t.Log(resp.String())
}

func TestResty2(t *testing.T) {
	resp, err := resty.New().R().Get("http://epms.istrongcloud.net/EPMS/AjaxHandler/DataHandler.ashx?MethodName=searchWeekSummaryData&YEAR=2022%27&WEEKS=7&U_ID=&DEPID=133&WEEKS0=0")
	assert.Equal(t, nil, err)
	t.Log(resp.String())
}

func TestResty3(t *testing.T) {
	// 不校验https证书，如果校验，类似企业微信认证可能出现错误：x509: certificate signed by unknown authority
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.R().Get("https://www.baidu.com/")
}

func TestResty4(t *testing.T) {
	// http://dmap.istrongcloud.net/charon/v1/agent/
	resp, err := resty.New().R().Get(`http://t1.tianditu.gov.cn/cia_c/wmts?SERVICE=WMTS&REQUEST=GetTile&VERSION=1.0.0&LAYER=cia&STYLE=default&TILEMATRIXSET=c&TILEMATRIX=4&TILEROW=4&TILECOL=13&FORMAT=tiles&tk=16554181e1d8f9f3b82ce84fe953c164`)
	assert.Equal(t, nil, err)
	t.Log(resp.String())
}
