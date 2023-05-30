package fn

import (
	"fmt"
	"reflect"
	"regexp"
)

// Format
// @param str
// @param data
// @example Format("{a}",map[string]any{}{"a":1})
func Format(src string, data any) string {
	switch data.(type) {
	case map[string]any:
		return FormatByMap(src, data.(map[string]any))
	}
	return FormatByStruct(src, data)
}

// FormatString
// @param str
// @param data
// @example FormatString("{a}",map[string]any{}{"a":1})
func FormatByMap(src string, data map[string]any) string {
	// 匹配花括号内的字符串`{xxx}`，常用于字符串格式化替换
	return regexp.MustCompile(`{([^}]+)}`).ReplaceAllStringFunc(src, func(key string) string {
		// 这里的key包含括号
		val := data[key[1:len(key)-1]]
		if val == nil {
			return ""
		}
		return fmt.Sprint(val)
	})
}

// FormatString
// @param str
// @param data
// @example FormatString("{a}",sturct{A:1})
func FormatByStruct(src string, data any) string {
	// 匹配花括号内的字符串`{xxx}`，常用于字符串格式化替换
	rv := reflect.ValueOf(data)
	return regexp.MustCompile(`{([^}]+)}`).ReplaceAllStringFunc(src, func(key string) string {
		// 这里的key包含括号
		val := rv.FieldByName(key[1 : len(key)-1])
		if !val.IsValid() {
			return ""
		}
		return fmt.Sprint(val)
	})
}
