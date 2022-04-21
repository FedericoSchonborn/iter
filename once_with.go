package iter

type OnceWithIterator[T any] struct {
	fn   func() T
	done bool
}

func OnceWith[T any](fn func() T) *OnceWithIterator[T] {
	return &OnceWithIterator[T]{
		fn:   fn,
		done: false,
	}
}

func (owi *OnceWithIterator[T]) Next() (_ T, ok bool) {
	if owi.done {
		var zero T
		return zero, false
	}

	owi.done = true
	return owi.fn(), true
}
