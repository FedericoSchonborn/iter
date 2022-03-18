package iter

type repeatWithIter[T any] struct {
	fn func() T
}

func RepeatWith[T any](fn func() T) Iterator[T] {
	return &repeatWithIter[T]{
		fn: fn,
	}
}

func (rwi *repeatWithIter[T]) Next() (_ T, ok bool) {
	return rwi.fn(), true
}
