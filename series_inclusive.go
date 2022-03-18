package iter

import "golang.org/x/exp/constraints"

type seriesInclusiveIter[T constraints.Integer] struct {
	start T
	end   T
	done  bool
}

func SeriesInclusive[T constraints.Integer](start, end T) Iterator[T] {
	return &seriesInclusiveIter[T]{
		start: start,
		end:   end,
		done:  false,
	}
}

func (sii *seriesInclusiveIter[T]) Next() (_ T, ok bool) {
	if sii.done || sii.start > sii.end {
		var zero T
		return zero, false
	}

	if sii.start < sii.end {
		start := sii.start
		sii.start = start + T(1)
		return start, true
	}

	sii.done = true
	return sii.start, true
}
