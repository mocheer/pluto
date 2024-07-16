package object

//
func Clone[T any](src T) T {
	// srcValue := reflect.ValueOf(src)
	// srcType := reflect.TypeOf(src)

	// destValue := reflect.New(srcType).Elem()

	// for i := 0; i < srcType.NumField(); i++ {
	// 	srcField := srcValue.Field(i)
	// 	destField := destValue.Field(i)

	// 	if srcField.CanSet() && destField.CanSet() {
	// 		destField.Set(srcField)
	// 	}
	// }

	// return destValue.Interface().(T)
	return src
}

// ShallowClone
// 浅拷贝
// T 如果只是结构体实例，不是指针，那就能直接复制
// 结构体的指针也会导致拷贝
func ShallowClone[T any](src T) T {
	return src
}
