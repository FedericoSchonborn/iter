package iter_test

import (
	"testing"

	"github.com/FedericoSchonborn/iter"
	"github.com/FedericoSchonborn/iter/slices"
)

func assert(t *testing.T, expr bool) {
	if !expr {
		t.Fatal("Assertion failed")
	}
}

func assertEqual[T comparable](t *testing.T, actual T, expected T) {
	if actual != expected {
		t.Fatalf("Expected value `%v`, got `%v` instead", expected, actual)
	}
}

func assertNext[T comparable, I iter.Iterator[T]](t *testing.T, iter I, expected T) {
	next, ok := iter.Next()
	if !ok {
		t.Fatal("Expected iterator to return a value")
	}

	assertEqual(t, next, expected)
}

func assertPeek[T comparable, I iter.Iterator[T]](t *testing.T, iter *iter.Peekable[T, I], expected T) {
	next, ok := iter.Peek()
	if !ok {
		t.Fatal("Expected iterator to return a value")
	}

	assertEqual(t, next, expected)
}

func assertNone[T any, I iter.Iterator[T]](t *testing.T, iter I) {
	if _, ok := iter.Next(); ok {
		t.Fatal("Expected iterator to return no value")
	}
}

func TestIterator(t *testing.T) {
	a := []int{1, 2, 3}
	it := slices.NewIterator(a)

	assertNext(t, it, 1)
	assertNext(t, it, 2)
	assertNext(t, it, 3)
	assertNone[int](t, it)
	assertNone[int](t, it)
	assertNone[int](t, it)
}

func TestAdvanceBy(t *testing.T) {
	a := []int{1, 2, 3, 4}
	it := slices.NewIterator(a)

	n, ok := iter.AdvanceBy[int](it, 2)
	assertEqual(t, n, 0)
	assertEqual(t, ok, true)

	assertNext(t, it, 3)

	n, ok = iter.AdvanceBy[int](it, 0)
	assertEqual(t, n, 0)
	assertEqual(t, ok, true)

	n, ok = iter.AdvanceBy[int](it, 100)
	assertEqual(t, n, 1)
	assertEqual(t, ok, false)
}

func TestAll_Basic(t *testing.T) {
	a := []int{1, 2, 3}
	assert(t, iter.All(slices.NewIterator(a), func(x int) bool { return x > 0 }))
	assert(t, !iter.All(slices.NewIterator(a), func(x int) bool { return x > 2 }))
}

func TestAll_FirstFalse(t *testing.T) {
	a := []int{1, 2, 3}
	it := slices.NewIterator(a)

	assert(t, !iter.All(it, func(x int) bool { return x != 2 }))

	assertNext(t, it, 3)
}

func TestPeekable(t *testing.T) {
	xs := []int{1, 2, 3}
	it := iter.NewPeekable[int](slices.NewIterator(xs))

	assertPeek(t, it, 1)
	assertNext(t, it, 1)
	assertNext(t, it, 2)
	assertPeek(t, it, 3)
	assertPeek(t, it, 3)
	assertNext(t, it, 3)
	assertNone[int](t, it)
	assertNone[int](t, it)
}
