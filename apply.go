package iter

type Apply[T, B any, I Iterator[T]] struct {
	iter I
	fn   func(T) B
}

func NewApply[T, B any, I Iterator[T]](iter I, fn func(value T) B) *Apply[T, B, I] {
	return &Apply[T, B, I]{
		iter: iter,
		fn:   fn,
	}
}

func (a *Apply[T, B, I]) Next() (_ B, ok bool) {
	next, ok := a.iter.Next()
	if !ok {
		return Zero[B](), false
	}

	return a.fn(next), true
}
