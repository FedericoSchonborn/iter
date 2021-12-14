package iter

type EmptyIter[Item any] struct{}

func Empty[Item any]() Iterator[Item] {
	return &EmptyIter[Item]{}
}

func (*EmptyIter[Item]) Next() (_ Item, ok bool) {
	var zero Item
	return zero, false
}
