package iter

import "golang.org/x/exp/constraints"

type SeriesIterator[T constraints.Integer] struct {
	start T
	end   T
}

func Series[T constraints.Integer](start, end T) *SeriesIterator[T] {
	return &SeriesIterator[T]{
		start: start,
		end:   end,
	}
}

func (si *SeriesIterator[T]) Next() (_ T, ok bool) {
	if si.start < si.end {
		start := si.start
		si.start = start + T(1)
		return start, true
	}

	var zero T
	return zero, false
}
