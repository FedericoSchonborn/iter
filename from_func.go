package iter

type FromFunc[T any] struct {
	fn func() (T, bool)
}

func NewFromFunc[T any](fn func() (T, bool)) *FromFunc[T] {
	return &FromFunc[T]{
		fn: fn,
	}
}

func (ff *FromFunc[T]) Next() (_ T, ok bool) {
	return ff.fn()
}
