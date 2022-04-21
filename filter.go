package iter

type Filter[T any, I Iterator[T]] struct {
	iter I
	fn   func(T) bool
}

func NewFilter[T any, I Iterator[T]](iter I, fn func(value T) bool) *Filter[T, I] {
	return &Filter[T, I]{
		iter: iter,
		fn:   fn,
	}
}

func (f *Filter[T, I]) Next() (_ T, ok bool) {
	for {
		next, ok := f.iter.Next()
		if !ok {
			return Zero[T](), false
		}

		if f.fn(next) {
			return next, true
		}
	}
}
