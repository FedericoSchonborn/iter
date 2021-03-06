package iter

var _ SizedBilateralIterator[any] = (*Empty[any])(nil)

type Empty[T any] struct{}

func NewEmpty[T any]() *Empty[T] {
	return &Empty[T]{}
}

func (*Empty[T]) Next() (_ T, ok bool) {
	return Zero[T](), false
}

func (*Empty[T]) NextBack() (_ T, ok bool) {
	return Zero[T](), false
}

func (*Empty[T]) Len() int {
	return 0
}
