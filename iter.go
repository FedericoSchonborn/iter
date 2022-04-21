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

func AdvanceBy[T any, I Iterator[T]](iter I, count int) (_ int, ok bool) {
	for i := 0; i < count; i++ {
		_, ok := iter.Next()
		if !ok {
			return i, false
		}
	}

	return 0, true
}

func All[T any, I Iterator[T]](iter I, fn func(value T) bool) bool {
	for {
		next, ok := iter.Next()
		if !ok {
			break
		}

		if !fn(next) {
			return false
		}
	}

	return true
}

func Any[T any, I Iterator[T]](iter I, fn func(value T) bool) bool {
	for {
		next, ok := iter.Next()
		if !ok {
			break
		}

		if fn(next) {
			return true
		}
	}

	return false
}

func Collect[T any, I Iterator[T]](iter I) []T {
	slice := []T{}
	for {
		next, ok := iter.Next()
		if !ok {
			break
		}

		slice = append(slice, next)
	}

	return slice
}

func Count[T any, I Iterator[T]](iter I) int {
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

func Fold[T, A any, I Iterator[T]](iter I, init A, fn func(acc A, value T) A) A {
	acc := init
	for {
		next, ok := iter.Next()
		if !ok {
			break
		}

		acc = fn(acc, next)
	}

	return acc
}

func ForEach[T any, I Iterator[T]](iter I, fn func(value T)) {
	for {
		next, ok := iter.Next()
		if !ok {
			break
		}

		fn(next)
	}
}

func Product[T constraints.Integer | constraints.Float | constraints.Complex, I Iterator[T]](iter I) T {
	init, ok := iter.Next()
	if !ok {
		return Zero[T]()
	}

	total := init
	for {
		next, ok := iter.Next()
		if !ok {
			break
		}

		total *= next
	}

	return total
}

func Sum[T constraints.Integer | constraints.Float | constraints.Complex | ~string, I Iterator[T]](iter I) T {
	var total T
	for {
		next, ok := iter.Next()
		if !ok {
			break
		}

		total += next
	}

	return total
}
