package iter

type OnceWith[T any] struct {
	fn   func() T
	done bool
}

func NewOnceWith[T any](fn func() T) *OnceWith[T] {
	return &OnceWith[T]{
		fn:   fn,
		done: false,
	}
}

func (ow *OnceWith[T]) Next() (_ T, ok bool) {
	if ow.done {
		var zero T
		return zero, false
	}

	ow.done = true
	return ow.fn(), true
}
