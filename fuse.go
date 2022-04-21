package iter

type Fuse[T any, I Iterator[T]] struct {
	iter  I
	fused bool
}

func NewFuse[T any, I Iterator[T]](iter I) *Fuse[T, I] {
	return &Fuse[T, I]{
		iter: iter,
	}
}

func (f *Fuse[T, I]) Next() (_ T, ok bool) {
	item, ok := f.iter.Next()
	if !ok {
		var zero T
		return zero, false
	}

	if f.fused {
		var zero T
		return zero, false
	}

	return item, true
}
