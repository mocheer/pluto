package fn

// GOEXPERIMENT=aliastypeparams
// type ResultError[T any] = struct {
// 	Value T
// 	Error error
// }

// func Result[T any](value any, err error) ResultError[T] {
// 	return ResultError{Value: value, Error: err}
// }

// func (m ResultError[T]) UnWrap() {
// 	if m.Error != nil {
// 		panic(m.Error)
// 	}
// }
