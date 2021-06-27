package JSON

import "github.com/tidwall/gjson"

// Parse
func Parse(data string) gjson.Result {
	return gjson.Parse(data)
}
