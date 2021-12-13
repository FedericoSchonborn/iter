package iter

type Iterator[Item any] interface {
	Next() (_ Item, ok bool)
}

func From[T any](a ...T) Iterator[T] {
	return FromSlice(a)
}

func AdvanceBy[Item any](iter Iterator[Item], n int) (_ int, ok bool) {
	for i := 0; i < n; i++ {
		_, ok := iter.Next()
		if !ok {
			return i, false
		}
	}

	return 0, true
}

func All[Item any](iter Iterator[Item], fn func(Item) bool) bool {
	for {
		item, ok := iter.Next()
		if !ok {
			break
		}

		if !fn(item) {
			return false
		}
	}

	return true
}

func Any[Item any](iter Iterator[Item], fn func(Item) bool) bool {
	for {
		item, ok := iter.Next()
		if !ok {
			break
		}

		if fn(item) {
			return true
		}
	}

	return false
}

func Count[Item any](iter Iterator[Item]) int {
	var n int
	for {
		_, ok := iter.Next()
		if !ok {
			break
		}

		n++
	}

	return n
}

func Fold[Item, B any](iter Iterator[Item], init B, fn func(B, Item) B) B {
	acc := init
	for {
		item, ok := iter.Next()
		if !ok {
			break
		}

		acc = fn(acc, item)
	}

	return acc
}

func ForEach[Item any](iter Iterator[Item], fn func(Item)) {
	for {
		item, ok := iter.Next()
		if !ok {
			break
		}

		fn(item)
	}
}

func IntoSlice[Item any](iter Iterator[Item]) []Item {
	slice := []Item{}
	for {
		item, ok := iter.Next()
		if !ok {
			break
		}

		slice = append(slice, item)
	}

	return slice
}
