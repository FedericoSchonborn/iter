package iter

type FromFunc[T any] struct {
	next func() (_ T, ok bool)
}

func NewFromFunc[T any](next func() (_ T, ok bool)) *FromFunc[T] {
	return &FromFunc[T]{
		next: next,
	}
}

func (ff *FromFunc[T]) Next() (_ T, ok bool) {
	return ff.next()
}
