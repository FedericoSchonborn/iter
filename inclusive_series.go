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
	if is.done {
		return Zero[T](), false
	}

	next := is.start
	if is.start >= is.end {
		is.done = true
		return next, true
	}

	is.start = next + T(1)
	return next, true
}
