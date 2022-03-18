package iter

type filterIter[T any] struct {
	iter Iterator[T]
	fn   func(T) bool
}

func Filter[T any](iter Iterator[T], fn func(T) bool) Iterator[T] {
	return &filterIter[T]{
		iter: iter,
		fn:   fn,
	}
}

func (fi *filterIter[T]) Next() (_ T, ok bool) {
	for {
		item, ok := fi.iter.Next()
		if !ok {
			var zero T
			return zero, false
		}

		if fi.fn(item) {
			return item, true
		}
	}
}
