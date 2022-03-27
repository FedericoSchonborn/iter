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

func (p *PeekableIterator[T, I]) Peek() (_ T, ok bool) {
	if p.peeked {
		return p.current, true
	}

	value, ok := p.inner.Next()
	if !ok {
		var zero T
		return zero, false
	}

	p.current = value
	p.peeked = true
	return value, true
}

func (p *PeekableIterator[T, I]) Next() (_ T, ok bool) {
	if p.peeked {
		var zero T

		value := p.current
		p.current = zero
		p.peeked = false
		return value, true
	}

	return p.inner.Next()
}
