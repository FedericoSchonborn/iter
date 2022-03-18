package iter_test

import (
	"testing"

	"github.com/FedericoSchonborn/go-iter"
)

func assert(t *testing.T, expr bool) {
	if !expr {
		t.Fatal("Assertion failed")
	}
}

func assertEqual[T comparable](t *testing.T, actual T, expected T) {
	if actual != expected {
		t.Fatalf("Expected value `%v`, got `%v` instead", actual, expected)
	}
}

func assertNext[T comparable](t *testing.T, iter iter.Iterator[T], expected T) {
	item, ok := iter.Next()
	if !ok {
		t.Fatal("Expected iterator to return a value")
	}

	assertEqual(t, item, expected)
}

func assertNone[T any](t *testing.T, iter iter.Iterator[T]) {
	if _, ok := iter.Next(); ok {
		t.Fatal("Expected iterator to return no value")
	}
}

func TestIterator(t *testing.T) {
	a := []int{1, 2, 3}
	it := iter.FromSlice(a)

	assertNext(t, it, 1)
	assertNext(t, it, 2)
	assertNext(t, it, 3)
	assertNone(t, it)
	assertNone(t, it)
	assertNone(t, it)
}

func TestAdvanceBy(t *testing.T) {
	a := []int{1, 2, 3, 4}
	it := iter.FromSlice(a)

	n, ok := iter.AdvanceBy(it, 2)
	assertEqual(t, n, 0)
	assertEqual(t, ok, true)

	assertNext(t, it, 3)

	n, ok = iter.AdvanceBy(it, 0)
	assertEqual(t, n, 0)
	assertEqual(t, ok, true)

	n, ok = iter.AdvanceBy(it, 100)
	assertEqual(t, n, 1)
	assertEqual(t, ok, false)
}

func TestAll_Basic(t *testing.T) {
	a := []int{1, 2, 3}
	assert(t, iter.All(iter.FromSlice(a), func(x int) bool { return x > 0 }))
	assert(t, !iter.All(iter.FromSlice(a), func(x int) bool { return x > 2 }))
}

func TestAll_FirstFalse(t *testing.T) {
	a := []int{1, 2, 3}
	it := iter.FromSlice(a)

	assert(t, !iter.All(it, func(x int) bool { return x != 2 }))

	assertNext(t, it, 3)
}

/*
func TestAny(t *testing.T) {
	greaterThan := func(n int) func(int) bool {
		return func(x int) bool {
			return x > n
		}
	}

	fmt.Println(iter.Any(iter.New(1, 2, 3), greaterThan(0)))
	fmt.Println(iter.Any(iter.New(1, 2, 3), greaterThan(5)))
	// Output:
	// true
	// false
}

func TestCount(t *testing.T) {
	a := iter.New(1, 2, 3)
	assertEqual(t, iter.Count(a), 3)

	b := iter.New(1, 2, 3, 4, 5)
	assertEqual(t, iter.Count(b), 5)
}

func TestEmpty(t *testing.T) {
	nope := iter.Empty[int]()
	fmt.Println(nope.Next())

	// Output:
	// 0 false
}

func TestEnumerate(t *testing.T) {
	iter := iter.Enumerate(iter.New('a', 'b', 'c'))
	PrintNext(iter, 4)

	// Output:
	// {0 97} true
	// {1 98} true
	// {2 99} true
	// {0 0} false
}

func TestFilter(t *testing.T) {
	isPositive := func(i int) bool {
		return i > 0
	}

	iter := iter.Filter(iter.New(0, 1, 2), isPositive)
	PrintNext(iter, 3)

	// Output:
	// 1 true
	// 2 true
	// 0 false
}

func TestFold(t *testing.T) {
	sum := iter.Fold(iter.New(1, 2, 3), 0, func(acc int, x int) int {
		return acc + x
	})

	fmt.Println(sum)
	// Output:
	// 6
}

func TestForEach(t *testing.T) {
	iter.ForEach(iter.New(1, 2, 3), func(x int) {
		fmt.Println(x * 2)
	})

	// Output:
	// 2
	// 4
	// 6
}

func TestFromFunc(t *testing.T) {
	count := 0
	counter := iter.FromFunc(func() (int, bool) {
		count++

		if count < 6 {
			return count, true
		}

		return 0, false
	})

	fmt.Println(iter.Collect(counter))
	// Output:
	// [1 2 3 4 5]
}

func TestCollect(t *testing.T) {
	double := func(i int) int {
		return i * 2
	}

	fmt.Println(iter.Collect(iter.Map(iter.New(1, 2, 3), double)))
	// Output:
	// [2 4 6]
}

func TestMap(t *testing.T) {
	double := func(i int) int {
		return i * 2
	}

	iter := iter.Map(iter.New(1, 2, 3), double)
	PrintNext(iter, 3)

	// Output:
	// 2 true
	// 4 true
	// 6 true
}

func TestOnce(t *testing.T) {
	one := iter.Once(1)
	PrintNext(one, 2)

	// Output:
	// 1 true
	// 0 false
}

func TestOnceWith(t *testing.T) {
	one := iter.OnceWith(func() int {
		return 1
	})
	PrintNext(one, 2)

	// Output:
	// 1 true
	// 0 false
}

func TestProduct(t *testing.T) {
	factorial := func(n int) int {
		return iter.Product(iter.SeriesInclusive(1, n))
	}

	fmt.Println(factorial(0))
	fmt.Println(factorial(1))
	fmt.Println(factorial(5))
	// Output:
	// 0
	// 1
	// 120
}

func TestRepeat(t *testing.T) {
	fourFours := iter.Take(iter.Repeat(4), 4)
	PrintNext(fourFours, 5)

	// Output:
	// 4 true
	// 4 true
	// 4 true
	// 4 true
	// 0 false
}

func TestRepeatWith(t *testing.T) {
	curr := 1
	pow2 := iter.Take(iter.RepeatWith(func() int {
		tmp := curr
		curr *= 2
		return tmp
	}), 4)
	PrintNext(pow2, 5)

	// Output:
	// 1 true
	// 2 true
	// 4 true
	// 8 true
	// 0 false
}

func TestSuccessors(t *testing.T) {
	powersOf10 := iter.Successors(1, func(n uint) (uint, bool) {
		_, lo := bits.Mul(n, 10)
		if lo >= math.MaxUint16 {
			return 0, false
		}

		return lo, true
	})
	fmt.Println(iter.Collect(powersOf10))

	// Output:
	// [1 10 100 1000 10000]
}

func TestSum(t *testing.T) {
	fmt.Println(iter.Sum(iter.New(1, 2, 3)))
	// Output:
	// 6
}

func TestTake(t *testing.T) {
	iter := iter.Take(iter.New(1, 2, 3), 2)
	PrintNext(iter, 3)

	// Output:
	// 1 true
	// 2 true
	// 0 false
}

func TestZip(t *testing.T) {
	iter := iter.Zip(iter.New(1, 2, 3), iter.New(4, 5, 6))
	PrintNext(iter, 4)

	// Output:
	// {1 4} true
	// {2 5} true
	// {3 6} true
	// {0 0} false
}
*/
