package iter

import "constraints"

type SeriesInclusiveIter[Item constraints.Integer] struct {
	start     Item
	end       Item
	exhausted bool
}

func SeriesInclusive[Item constraints.Integer](start, end Item) Iterator[Item] {
	return &SeriesInclusiveIter[Item]{
		start:     start,
		end:       end,
		exhausted: false,
	}
}

func (sii *SeriesInclusiveIter[Item]) Next() (_ Item, ok bool) {
	if sii.exhausted || sii.start > sii.end {
		var zero Item
		return zero, false
	}

	if sii.start < sii.end {
		n := sii.start + Item(1)
		start := sii.start
		sii.start = n
		return start, true
	}

	sii.exhausted = true
	return sii.start, true
}
