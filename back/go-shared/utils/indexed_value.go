package utils

type IndexedValue[T any] struct {
	Index int
	Value T
}

func NewIndexedValue[T any](index int, value T) IndexedValue[T] {
	return IndexedValue[T]{Index: index, Value: value}
}
