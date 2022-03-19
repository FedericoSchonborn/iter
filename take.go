package iter

type TakeIterator[T any] struct {
	iter  Iterator[T]
	count int
}

func Take[T any](iter Iterator[T], n int) *TakeIterator[T] {
	return &TakeIterator[T]{
		iter:  iter,
		count: n,
	}
}

func (ti *TakeIterator[T]) Next() (_ T, ok bool) {
	if ti.count <= 0 {
		var zero T
		return zero, false
	}

	item, ok := ti.iter.Next()
	if !ok {
		var zero T
		return zero, false
	}

	ti.count--
	return item, true
}
