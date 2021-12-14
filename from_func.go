package iter

type FromFuncIter[Item any] struct {
	fn func() (Item, bool)
}

func FromFunc[Item any](fn func() (Item, bool)) Iterator[Item] {
	return &FromFuncIter[Item]{
		fn: fn,
	}
}

func (ffi *FromFuncIter[Item]) Next() (_ Item, ok bool) {
	return ffi.fn()
}
