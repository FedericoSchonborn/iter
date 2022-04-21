package iter

type MapIterator[T, B any, I Iterator[T]] struct {
	iter I
	fn   func(T) B
}

func Map[T, B any, I Iterator[T]](iter I, fn func(T) B) *MapIterator[T, B, I] {
	return &MapIterator[T, B, I]{
		iter: iter,
		fn:   fn,
	}
}

func (mi *MapIterator[T, B, I]) Next() (_ B, ok bool) {
	item, ok := mi.iter.Next()
	if !ok {
		var zero B
		return zero, false
	}

	return mi.fn(item), true
}
