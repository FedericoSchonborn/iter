package iter

type TakeIterator[T any, I Iterator[T]] struct {
	iter  I
	count int
}

func Take[T any, I Iterator[T]](iter I, n int) *TakeIterator[T, I] {
	return &TakeIterator[T, I]{
		iter:  iter,
		count: n,
	}
}

func (ti *TakeIterator[T, I]) Next() (_ T, ok bool) {
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
