package iter

type Map[T, B any, I Iterator[T]] struct {
	iter I
	fn   func(T) B
}

func NewMap[T, B any, I Iterator[T]](iter I, fn func(value T) B) *Map[T, B, I] {
	return &Map[T, B, I]{
		iter: iter,
		fn:   fn,
	}
}

func (m *Map[T, B, I]) Next() (_ B, ok bool) {
	next, ok := m.iter.Next()
	if !ok {
		return Zero[B](), false
	}

	return m.fn(next), true
}
