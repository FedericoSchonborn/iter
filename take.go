package iter

type Take[T any, I Iterator[T]] struct {
	iter  I
	count int
}

func NewTake[T any, I Iterator[T]](iter I, n int) *Take[T, I] {
	return &Take[T, I]{
		iter:  iter,
		count: n,
	}
}

func (t *Take[T, I]) Next() (_ T, ok bool) {
	if t.count <= 0 {
		var zero T
		return zero, false
	}

	item, ok := t.iter.Next()
	if !ok {
		var zero T
		return zero, false
	}

	t.count--
	return item, true
}
