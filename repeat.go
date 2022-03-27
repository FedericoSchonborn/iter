package iter

type Repeat[T any] struct {
	item T
}

func NewRepeat[T any](item T) *Repeat[T] {
	return &Repeat[T]{
		item: item,
	}
}

func (r *Repeat[T]) Next() (_ T, ok bool) {
	return r.item, true
}
