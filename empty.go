package iter

var _ SizedBilateralIterator[any] = (*EmptyIterator[any])(nil)

type EmptyIterator[T any] struct{}

func Empty[T any]() *EmptyIterator[T] {
	return &EmptyIterator[T]{}
}

func (*EmptyIterator[T]) Next() (_ T, ok bool) {
	var zero T
	return zero, false
}

func (*EmptyIterator[T]) NextBack() (_ T, ok bool) {
	var zero T
	return zero, false
}

func (*EmptyIterator[T]) Len() int {
	return 0
}
