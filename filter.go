package iter

type FilterIter[Item any] struct {
	iter Iterator[Item]
	fn   func(Item) bool
}

func Filter[Item any](iter Iterator[Item], fn func(Item) bool) Iterator[Item] {
	return &FilterIter[Item]{
		iter: iter,
		fn:   fn,
	}
}

func (fi *FilterIter[Item]) Next() (_ Item, ok bool) {
	for {
		item, ok := fi.iter.Next()
		if !ok {
			var zero Item
			return zero, false
		}

		if fi.fn(item) {
			return item, true
		}
	}
}
