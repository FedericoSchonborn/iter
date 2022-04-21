package iter

type RepeatWithIterator[T any] struct {
	fn func() T
}

func RepeatWith[T any](fn func() T) *RepeatWithIterator[T] {
	return &RepeatWithIterator[T]{
		fn: fn,
	}
}

func (rwi *RepeatWithIterator[T]) Next() (_ T, ok bool) {
	return rwi.fn(), true
}
