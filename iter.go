package iter

type Iterator[T any] interface {
	Next() (_ T, ok bool)
}
