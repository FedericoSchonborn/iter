// Package iter provides generic lazy iterators compatible with Go 1.18+.
package iter

import "golang.org/x/exp/constraints"

type Iterator[T any] interface {
	Next() (_ T, ok bool)
}

type BilateralIterator[T any] interface {
	Iterator[T]
	NextBack() (_ T, ok bool)
}

type SizedIterator[T any] interface {
	Iterator[T]
	Len() int
}

type SizedBilateralIterator[T any] interface {
	BilateralIterator[T]
	SizedIterator[T]
}

func AdvanceBy[T any](iter Iterator[T], n int) (_ int, ok bool) {
	for i := 0; i < n; i++ {
		_, ok := iter.Next()
		if !ok {
			return i, false
		}
	}

	return 0, true
}

func All[T any](iter Iterator[T], fn func(T) bool) bool {
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

func Any[T any](iter Iterator[T], fn func(T) bool) bool {
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

func Collect[T any](iter Iterator[T]) []T {
	slice := []T{}
	for {
		item, ok := iter.Next()
		if !ok {
			break
		}

		slice = append(slice, item)
	}

	return slice
}

func Count[T any](iter Iterator[T]) int {
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

func Fold[T, B any](iter Iterator[T], init B, fn func(B, T) B) B {
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

func ForEach[T any](iter Iterator[T], fn func(T)) {
	for {
		item, ok := iter.Next()
		if !ok {
			break
		}

		fn(item)
	}
}

func Product[T constraints.Integer | constraints.Float | constraints.Complex](iter Iterator[T]) T {
	total, ok := iter.Next()
	if !ok {
		var zero T
		return zero
	}

	for {
		item, ok := iter.Next()
		if !ok {
			break
		}

		total *= item
	}

	return total
}

func Sum[T constraints.Integer | constraints.Float | constraints.Complex | ~string](iter Iterator[T]) T {
	var total T
	for {
		item, ok := iter.Next()
		if !ok {
			break
		}

		total += item
	}

	return total
}
