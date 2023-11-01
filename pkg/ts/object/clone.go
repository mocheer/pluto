package object

import "reflect"

func Clone[T any](src T) T {
	srcValue := reflect.ValueOf(src)
	srcType := reflect.TypeOf(src)

	destValue := reflect.New(srcType).Elem()

	for i := 0; i < srcType.NumField(); i++ {
		srcField := srcValue.Field(i)
		destField := destValue.Field(i)

		if srcField.CanSet() && destField.CanSet() {
			destField.Set(srcField)
		}
	}

	return destValue.Interface().(T)
}
