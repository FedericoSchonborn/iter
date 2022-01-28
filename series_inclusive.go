package iter

import "constraints"

type SeriesInclusiveIter[Item constraints.Integer] struct {
	start Item
	end   Item
	done  bool
}

func SeriesInclusive[Item constraints.Integer](start, end Item) Iterator[Item] {
	return &SeriesInclusiveIter[Item]{
		start: start,
		end:   end,
		done:  false,
	}
}

func (sii *SeriesInclusiveIter[Item]) Next() (_ Item, ok bool) {
	if sii.done || sii.start > sii.end {
		var zero Item
		return zero, false
	}

	if sii.start < sii.end {
		start := sii.start
		sii.start = start + Item(1)
		return start, true
	}

	sii.done = true
	return sii.start, true
}
