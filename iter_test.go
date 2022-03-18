package iter_test

import (
	"fmt"
	"math"
	"math/bits"

	"github.com/FedericoSchonborn/go-iter"
)

func PrintNext[T any](iter iter.Iterator[T], n int) {
	for i := 0; i < n; i++ {
		fmt.Println(iter.Next())
	}
}

func ExampleIterator() {
	iter := iter.New(1, 2, 3)
	PrintNext(iter, 6)
	// Output:
	// 1 true
	// 2 true
	// 3 true
	// 0 false
	// 0 false
	// 0 false
}

func ExampleAdvanceBy() {
	var ok bool

	i := iter.New(1, 2, 3, 4)

	_, ok = iter.AdvanceBy(i, 2)
	fmt.Println(ok)

	item, _ := i.Next()
	fmt.Println(item)

	_, ok = iter.AdvanceBy(i, 0)
	fmt.Println(ok)

	n, _ := iter.AdvanceBy(i, 100)
	fmt.Println(n)

	// Output:
	// true
	// 3
	// true
	// 1
}

func ExampleAll() {
	greaterThan := func(n int) func(int) bool {
		return func(x int) bool {
			return x > n
		}
	}

	fmt.Println(iter.All(iter.New(1, 2, 3), greaterThan(0)))
	fmt.Println(iter.All(iter.New(1, 2, 3), greaterThan(2)))
	// Output:
	// true
	// false
}

func ExampleAny() {
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

func ExampleCount() {
	fmt.Println(iter.Count(iter.New(1, 2, 3)))
	fmt.Println(iter.Count(iter.New(1, 2, 3, 4, 5)))
	// Output:
	// 3
	// 5
}

func ExampleEmpty() {
	nope := iter.Empty[int]()
	fmt.Println(nope.Next())

	// Output:
	// 0 false
}

func ExampleEnumerate() {
	iter := iter.Enumerate(iter.New('a', 'b', 'c'))
	PrintNext(iter, 4)

	// Output:
	// {0 97} true
	// {1 98} true
	// {2 99} true
	// {0 0} false
}

func ExampleFilter() {
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

func ExampleFold() {
	sum := iter.Fold(iter.New(1, 2, 3), 0, func(acc int, x int) int {
		return acc + x
	})

	fmt.Println(sum)
	// Output:
	// 6
}

func ExampleForEach() {
	iter.ForEach(iter.New(1, 2, 3), func(x int) {
		fmt.Println(x * 2)
	})

	// Output:
	// 2
	// 4
	// 6
}

func ExampleFromFunc() {
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

func ExampleCollect() {
	double := func(i int) int {
		return i * 2
	}

	fmt.Println(iter.Collect(iter.Map(iter.New(1, 2, 3), double)))
	// Output:
	// [2 4 6]
}

func ExampleMap() {
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

func ExampleOnce() {
	one := iter.Once(1)
	PrintNext(one, 2)

	// Output:
	// 1 true
	// 0 false
}

func ExampleOnceWith() {
	one := iter.OnceWith(func() int {
		return 1
	})
	PrintNext(one, 2)

	// Output:
	// 1 true
	// 0 false
}

func ExampleProduct() {
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

func ExampleRepeat() {
	fourFours := iter.Take(iter.Repeat(4), 4)
	PrintNext(fourFours, 5)

	// Output:
	// 4 true
	// 4 true
	// 4 true
	// 4 true
	// 0 false
}

func ExampleRepeatWith() {
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

func ExampleSuccessors() {
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

func ExampleSum() {
	fmt.Println(iter.Sum(iter.New(1, 2, 3)))
	// Output:
	// 6
}

func ExampleTake() {
	iter := iter.Take(iter.New(1, 2, 3), 2)
	PrintNext(iter, 3)

	// Output:
	// 1 true
	// 2 true
	// 0 false
}

func ExampleZip() {
	iter := iter.Zip(iter.New(1, 2, 3), iter.New(4, 5, 6))
	PrintNext(iter, 4)

	// Output:
	// {1 4} true
	// {2 5} true
	// {3 6} true
	// {0 0} false
}
