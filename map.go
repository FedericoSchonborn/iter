package iter

type MapIter[Item, B any] struct {
	iter Iterator[Item]
	fn   func(Item) B
}

func Map[Item, B any](iter Iterator[Item], fn func(Item) B) Iterator[B] {
	return &MapIter[Item, B]{
		iter: iter,
		fn:   fn,
	}
}

func (mi *MapIter[Item, B]) Next() (_ B, ok bool) {
	item, ok := mi.iter.Next()
	if !ok {
		var zero B
		return zero, false
	}

	return mi.fn(item), true
}
