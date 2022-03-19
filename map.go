package iter

type MapIterator[T, B any] struct {
	iter Iterator[T]
	fn   func(T) B
}

func Map[T, B any](iter Iterator[T], fn func(T) B) *MapIterator[T, B] {
	return &MapIterator[T, B]{
		iter: iter,
		fn:   fn,
	}
}

func (mi *MapIterator[T, B]) Next() (_ B, ok bool) {
	item, ok := mi.iter.Next()
	if !ok {
		var zero B
		return zero, false
	}

	return mi.fn(item), true
}
