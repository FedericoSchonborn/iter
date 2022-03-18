package iter

type fromFuncIter[T any] struct {
	fn func() (T, bool)
}

func FromFunc[T any](fn func() (T, bool)) Iterator[T] {
	return &fromFuncIter[T]{
		fn: fn,
	}
}

func (ffi *fromFuncIter[T]) Next() (_ T, ok bool) {
	return ffi.fn()
}
