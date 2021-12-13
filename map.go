package iter

type MapFunc[Item, B any] func(Item) B

type MapIter[I Iterator[Item], Item, B any] struct {
	iter I
	fn   MapFunc[Item, B]
}

func Map[I Iterator[Item], Item, B any](iter I, fn MapFunc[Item, B]) Iterator[B] {
	return &MapIter[I, Item, B]{
		iter: iter,
		fn:   fn,
	}
}

func (mi *MapIter[I, Item, B]) Next() (_ B, ok bool) {
	value, ok := mi.iter.Next()
	if !ok {
		var zero B
		return zero, false
	}

	return mi.fn(value), true
}
