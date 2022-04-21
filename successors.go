package iter

type Successors[T any] struct {
	first   bool
	current T
	fn      func(T) (T, bool)
}

func NewSuccessors[T any](first T, fn func(T) (T, bool)) *Successors[T] {
	return &Successors[T]{
		first:   true,
		current: first,
		fn:      fn,
	}
}

func (s *Successors[T]) Next() (_ T, ok bool) {
	if s.first {
		s.first = false
		return s.current, true
	}

	s.current, ok = s.fn(s.current)
	if !ok {
		var zero T
		return zero, false
	}

	return s.current, true
}
