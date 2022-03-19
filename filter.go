package iter

type FilterIterator[T any, I Iterator[T]] struct {
	iter I
	fn   func(T) bool
}

func Filter[T any, I Iterator[T]](iter I, fn func(T) bool) *FilterIterator[T, I] {
	return &FilterIterator[T, I]{
		iter: iter,
		fn:   fn,
	}
}

func (fi *FilterIterator[T, I]) Next() (_ T, ok bool) {
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
