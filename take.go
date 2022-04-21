package iter

type Take[T any, I Iterator[T]] struct {
	iter  I
	count int
}

func NewTake[T any, I Iterator[T]](iter I, count int) *Take[T, I] {
	return &Take[T, I]{
		iter:  iter,
		count: count,
	}
}

func (t *Take[T, I]) Next() (_ T, ok bool) {
	if t.count <= 0 {
		return Zero[T](), false
	}

	next, ok := t.iter.Next()
	if !ok {
		return Zero[T](), false
	}

	t.count--
	return next, true
}
