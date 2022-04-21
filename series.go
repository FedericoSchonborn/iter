package iter

import "golang.org/x/exp/constraints"

type Series[T constraints.Integer] struct {
	start T
	end   T
}

func NewSeries[T constraints.Integer](start, end T) *Series[T] {
	return &Series[T]{
		start: start,
		end:   end,
	}
}

func (s *Series[T]) Next() (_ T, ok bool) {
	if s.start < s.end {
		start := s.start
		s.start = start + T(1)
		return start, true
	}

	var zero T
	return zero, false
}
