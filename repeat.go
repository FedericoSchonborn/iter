package iter

type Repeat[T any] struct {
	value T
}

func NewRepeat[T any](value T) *Repeat[T] {
	return &Repeat[T]{
		value: value,
	}
}

func (r *Repeat[T]) Next() (_ T, ok bool) {
	return r.value, true
}
