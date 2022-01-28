package iter

import "constraints"

type RangeIter[Item constraints.Integer] struct {
	start Item
	end   Item
}

func Range[Item constraints.Integer](start, end Item) Iterator[Item] {
	return &RangeIter[Item]{
		start: start,
		end:   end,
	}
}

func (ri *RangeIter[Item]) Next() (_ Item, ok bool) {
	if ri.start < ri.end {
		n := ri.start + Item(1)
		start := ri.start
		ri.start = n
		return start, true
	}

	var zero Item
	return zero, false
}
