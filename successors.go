package iter

type Successors[T any] struct {
	first bool
	last  T
	next  func(T) (T, bool)
}

func NewSuccessors[T any](init T, next func(value T) (_ T, ok bool)) *Successors[T] {
	return &Successors[T]{
		first: true,
		last:  init,
		next:  next,
	}
}

func (s *Successors[T]) Next() (_ T, ok bool) {
	if s.first {
		s.first = false
		return s.last, true
	}

	next, ok := s.next(s.last)
	if !ok {
		return Zero[T](), false
	}

	s.last = next
	return next, true
}
