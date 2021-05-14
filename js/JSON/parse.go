package JSON

import "github.com/tidwall/gjson"

// Parse
func Parse(data string) interface{} {
	return gjson.Parse(data).Value()
}
