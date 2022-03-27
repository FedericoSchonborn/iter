package iter

type RepeatWith[T any] struct {
	fn func() T
}

func NewRepeatWith[T any](fn func() T) *RepeatWith[T] {
	return &RepeatWith[T]{
		fn: fn,
	}
}

func (rw *RepeatWith[T]) Next() (_ T, ok bool) {
	return rw.fn(), true
}
