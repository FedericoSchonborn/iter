package iter

type Fuse[T any, I Iterator[T]] struct {
	iter I
	done bool
}

func NewFuse[T any, I Iterator[T]](iter I) *Fuse[T, I] {
	return &Fuse[T, I]{
		iter: iter,
	}
}

func (f *Fuse[T, I]) Next() (_ T, ok bool) {
	if f.done {
		return Zero[T](), false
	}

	next, ok := f.iter.Next()
	if !ok {
		return Zero[T](), false
	}

	return next, true
}
