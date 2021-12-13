package iter

type MapFunc[T, U any] func(T) U

type MapIter[I Iterator[T], T, U any] struct {
	iter I
	fn   MapFunc[T, U]
}

func Map[I Iterator[T], T, U any](iter I, fn MapFunc[T, U]) Iterator[U] {
	return &MapIter[I, T, U]{
		iter: iter,
		fn:   fn,
	}
}

func (mi *MapIter[I, T, U]) Next() (_ U, ok bool) {
	value, ok := mi.iter.Next()
	if !ok {
		var zero U
		return zero, false
	}

	return mi.fn(value), true
}
