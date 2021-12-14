package iter

type RepeatWithIter[Item any] struct {
	fn func() Item
}

func RepeatWith[Item any](fn func() Item) Iterator[Item] {
	return &RepeatWithIter[Item]{
		fn: fn,
	}
}

func (rwi *RepeatWithIter[Item]) Next() (_ Item, ok bool) {
	return rwi.fn(), true
}
