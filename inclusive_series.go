package iter

import "golang.org/x/exp/constraints"

type InclusiveSeriesIterator[T constraints.Integer] struct {
	start T
	end   T
	done  bool
}

func InclusiveSeries[T constraints.Integer](start, end T) *InclusiveSeriesIterator[T] {
	return &InclusiveSeriesIterator[T]{
		start: start,
		end:   end,
		done:  false,
	}
}

func (isi *InclusiveSeriesIterator[T]) Next() (_ T, ok bool) {
	if isi.done || isi.start > isi.end {
		var zero T
		return zero, false
	}

	if isi.start < isi.end {
		start := isi.start
		isi.start = start + T(1)
		return start, true
	}

	isi.done = true
	return isi.start, true
}
