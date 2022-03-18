package iter

type takeIter[T any] struct {
	iter  Iterator[T]
	count int
}

func Take[T any](iter Iterator[T], n int) Iterator[T] {
	return &takeIter[T]{
		iter:  iter,
		count: n,
	}
}

func (ti *takeIter[T]) Next() (_ T, ok bool) {
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
