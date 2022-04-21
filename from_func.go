package iter

type FromFuncIterator[T any] struct {
	fn func() (T, bool)
}

func FromFunc[T any](fn func() (T, bool)) *FromFuncIterator[T] {
	return &FromFuncIterator[T]{
		fn: fn,
	}
}

func (ffi *FromFuncIterator[T]) Next() (_ T, ok bool) {
	return ffi.fn()
}
