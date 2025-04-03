package gstd

import "errors"

type result[T any] struct {
	Value T
	Error error
}

type Result[T any] = *result[T]

func NewResult[T any](value T, err error) Result[T] {
	return &result[T]{Value: value, Error: err}
}

func (m result[T]) UnWrap() T {
	return unWrap(m)
}

func (m *result[T]) Join(err error) Result[T] {
	m.Error = errors.Join(m.Error, err)
	return m
}

func unWrap[T any](r result[T]) T {
	if r.Error != nil {
		panic(r.Error)
	}
	return r.Value
}
