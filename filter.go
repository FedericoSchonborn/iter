package iter

type Filter[T any, I Iterator[T]] struct {
	iter I
	fn   func(T) bool
}

func NewFilter[T any, I Iterator[T]](iter I, fn func(T) bool) *Filter[T, I] {
	return &Filter[T, I]{
		iter: iter,
		fn:   fn,
	}
}

func (f *Filter[T, I]) Next() (_ T, ok bool) {
	for {
		item, ok := f.iter.Next()
		if !ok {
			var zero T
			return zero, false
		}

		if f.fn(item) {
			return item, true
		}
	}
}
