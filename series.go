package iter

import "constraints"

type SeriesIter[Item constraints.Integer] struct {
	start Item
	end   Item
}

func Series[Item constraints.Integer](start, end Item) Iterator[Item] {
	return &SeriesIter[Item]{
		start: start,
		end:   end,
	}
}

func (si *SeriesIter[Item]) Next() (_ Item, ok bool) {
	if si.start < si.end {
		n := si.start + Item(1)
		start := si.start
		si.start = n
		return start, true
	}

	var zero Item
	return zero, false
}
