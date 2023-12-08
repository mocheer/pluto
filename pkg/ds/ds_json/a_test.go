package ds_json_test

import (
	"testing"

	"github.com/tidwall/gjson"
)

func TestXxx(t *testing.T) {
	t.Log(gjson.Parse(`{}`).IsObject())
	t.Log(gjson.Parse(`[]`).IsObject())
	t.Log(gjson.Parse(`"{}"`).IsObject())
	t.Log(gjson.Parse(`string`).IsObject())
	//
	t.Log(gjson.Parse(`string`).Exists())
	t.Log(gjson.Parse(`string`).Value())
	t.Log(gjson.Parse(`"string"`).String())
	t.Log(gjson.Parse(`string`).String())
}
