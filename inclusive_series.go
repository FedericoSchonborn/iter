package iter

import "golang.org/x/exp/constraints"

type InclusiveSeries[T constraints.Integer] struct {
	start T
	end   T
	done  bool
}

func NewInclusiveSeries[T constraints.Integer](start, end T) *InclusiveSeries[T] {
	return &InclusiveSeries[T]{
		start: start,
		end:   end,
		done:  false,
	}
}

func (is *InclusiveSeries[T]) Next() (_ T, ok bool) {
	if is.done || is.start > is.end {
		var zero T
		return zero, false
	}

	if is.start < is.end {
		start := is.start
		is.start = start + T(1)
		return start, true
	}

	is.done = true
	return is.start, true
}
