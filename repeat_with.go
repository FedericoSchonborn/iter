package iter

type RepeatWith[T any] struct {
	next func() T
}

func NewRepeatWith[T any](next func() (value T)) *RepeatWith[T] {
	return &RepeatWith[T]{
		next: next,
	}
}

func (rw *RepeatWith[T]) Next() (_ T, ok bool) {
	return rw.next(), true
}
