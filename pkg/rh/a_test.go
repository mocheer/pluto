package rh_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/rh"
)

func TestGet(t *testing.T) {
	data, err := rh.Get("https://www.baidu.com/")
	t.Log(string(data))
	t.Log(err)
}

func TestGet2(t *testing.T) {
	data, err := rh.Get("http://epms.istrongcloud.net/EPMS/AjaxHandler/DataHandler.ashx?MethodName=searchWeekSummaryData&YEAR=2022%27&WEEKS=7&U_ID=&DEPID=133&WEEKS0=0")
	t.Log(string(data))
	t.Log(err)
	t.Error("错误")

}
