package iter

type OnceWith[T any] struct {
	next func() T
	done bool
}

func NewOnceWith[T any](next func() (value T)) *OnceWith[T] {
	return &OnceWith[T]{
		next: next,
		done: false,
	}
}

func (ow *OnceWith[T]) Next() (_ T, ok bool) {
	if ow.done {
		return Zero[T](), false
	}

	ow.done = true
	return ow.next(), true
}
