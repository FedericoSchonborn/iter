package iter

type PeekableIterator[T any, I Iterator[T]] struct {
	inner   I
	current T
	peeked  bool
}

func Peekable[T any, I Iterator[T]](inner I) *PeekableIterator[T, I] {
	return &PeekableIterator[T, I]{
		inner: inner,
	}
}

func (pi *PeekableIterator[T, I]) Peek() (_ T, ok bool) {
	if pi.peeked {
		return pi.current, true
	}

	value, ok := pi.inner.Next()
	if !ok {
		var zero T
		return zero, false
	}

	pi.current = value
	pi.peeked = true
	return value, true
}

func (pi *PeekableIterator[T, I]) Next() (_ T, ok bool) {
	if pi.peeked {
		var zero T

		value := pi.current
		pi.current = zero
		pi.peeked = false
		return value, true
	}

	return pi.inner.Next()
}
