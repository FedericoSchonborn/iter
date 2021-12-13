package iter

type Iterator[Item any] interface {
	Next() (_ Item, ok bool)
}
