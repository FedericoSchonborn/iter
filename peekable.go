package iter

type Peekable[T any, I Iterator[T]] struct {
	iter I
	last T
	peek bool
}

func NewPeekable[T any, I Iterator[T]](inner I) *Peekable[T, I] {
	return &Peekable[T, I]{
		iter: inner,
		peek: true,
	}
}

func (p *Peekable[T, I]) Peek() (_ T, ok bool) {
	if !p.peek {
		return p.last, true
	}

	next, ok := p.iter.Next()
	if !ok {
		return Zero[T](), false
	}

	p.last, p.peek = next, false
	return next, true
}

func (p *Peekable[T, I]) Next() (_ T, ok bool) {
	if !p.peek {
		next := p.last
		p.last, p.peek = Zero[T](), true
		return next, true
	}

	return p.iter.Next()
}
