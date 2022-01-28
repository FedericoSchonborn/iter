package iter

import "constraints"

type RangeInclusiveIter[Item constraints.Integer] struct {
	start     Item
	end       Item
	exhausted bool
}

func RangeInclusive[Item constraints.Integer](start, end Item) Iterator[Item] {
	return &RangeInclusiveIter[Item]{
		start:     start,
		end:       end,
		exhausted: false,
	}
}

func (rii *RangeInclusiveIter[Item]) Next() (_ Item, ok bool) {
	if rii.exhausted || rii.start > rii.end {
		var zero Item
		return zero, false
	}

	if rii.start < rii.end {
		n := rii.start + Item(1)
		start := rii.start
		rii.start = n
		return start, true
	}

	rii.exhausted = true
	return rii.start, true
}
