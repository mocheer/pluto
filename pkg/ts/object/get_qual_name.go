package object

import (
	"reflect"
	"strings"
)

func GetQualName(v any) string {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	s := t.PkgPath()
	return s[strings.LastIndex(s, "/")+1:] + "." + t.Name()
}
