package iter

type emptyIter[T any] struct{}

func Empty[T any]() Iterator[T] {
	return &emptyIter[T]{}
}

func (*emptyIter[T]) Next() (_ T, ok bool) {
	var zero T
	return zero, false
}
