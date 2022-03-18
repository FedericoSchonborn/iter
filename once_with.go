package iter

type onceWithIter[T any] struct {
	fn   func() T
	done bool
}

func OnceWith[T any](fn func() T) Iterator[T] {
	return &onceWithIter[T]{
		fn:   fn,
		done: false,
	}
}

func (owi *onceWithIter[T]) Next() (_ T, ok bool) {
	if owi.done {
		var zero T
		return zero, false
	}

	owi.done = true
	return owi.fn(), true
}
