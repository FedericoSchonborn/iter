package iter_test

import (
	"fmt"
	"math"
	"math/bits"

	. "github.com/fdschonborn/go-iter"
)

func greaterThan(n int) func(int) bool {
	return func(x int) bool {
		return x > n
	}
}

func isPositive(n int) bool {
	return n > 0
}

func double(n int) int {
	return n * 2
}

func printNextN[Item any](iter Iterator[Item], n int) {
	for i := 0; i < n; i++ {
		fmt.Println(iter.Next())
	}
}

func ExampleAdvanceBy() {
	var ok bool

	iter := New(1, 2, 3, 4)

	_, ok = AdvanceBy(iter, 2)
	fmt.Println(ok)

	item, _ := iter.Next()
	fmt.Println(item)

	_, ok = AdvanceBy(iter, 0)
	fmt.Println(ok)

	n, _ := AdvanceBy(iter, 100)
	fmt.Println(n)

	// Output:
	// true
	// 3
	// true
	// 1
}

func ExampleAll() {
	fmt.Println(All(New(1, 2, 3), greaterThan(0)))
	fmt.Println(All(New(1, 2, 3), greaterThan(2)))
	// Output:
	// true
	// false
}

func ExampleAny() {
	fmt.Println(Any(New(1, 2, 3), greaterThan(0)))
	fmt.Println(Any(New(1, 2, 3), greaterThan(5)))
	// Output:
	// true
	// false
}

func ExampleCount() {
	fmt.Println(Count(New(1, 2, 3)))
	fmt.Println(Count(New(1, 2, 3, 4, 5)))
	// Output:
	// 3
	// 5
}

func ExampleEmpty() {
	nope := Empty[int]()
	fmt.Println(nope.Next())

	// Output:
	// 0 false
}

func ExampleEnumerate() {
	iter := Enumerate(New('a', 'b', 'c'))
	printNextN(iter, 4)

	// Output:
	// {0 97} true
	// {1 98} true
	// {2 99} true
	// {0 0} false
}

func ExampleFilter() {
	iter := Filter(New(0, 1, 2), isPositive)
	printNextN(iter, 3)

	// Output:
	// 1 true
	// 2 true
	// 0 false
}

func ExampleFold() {
	sum := Fold(New(1, 2, 3), 0, func(acc int, x int) int {
		return acc + x
	})

	fmt.Println(sum)
	// Output:
	// 6
}

func ExampleForEach() {
	ForEach(New(1, 2, 3), func(x int) {
		fmt.Println(x * 2)
	})

	// Output:
	// 2
	// 4
	// 6
}

func ExampleFromFunc() {
	count := 0
	counter := FromFunc(func() (int, bool) {
		count++

		if count < 6 {
			return count, true
		}

		return 0, false
	})

	fmt.Println(Collect(counter))
	// Output:
	// [1 2 3 4 5]
}

func ExampleCollect() {
	fmt.Println(Collect(Map(New(1, 2, 3), double)))
	// Output:
	// [2 4 6]
}

func ExampleMap() {
	iter := Map(New(1, 2, 3), double)
	printNextN(iter, 3)

	// Output:
	// 2 true
	// 4 true
	// 6 true
}

func ExampleOnce() {
	one := Once(1)
	printNextN(one, 2)

	// Output:
	// 1 true
	// 0 false
}

func ExampleOnceWith() {
	one := OnceWith(func() int {
		return 1
	})
	printNextN(one, 2)

	// Output:
	// 1 true
	// 0 false
}

func ExampleProduct() {
	factorial := func(n int) int {
		return Product(SeriesInclusive(1, n))
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
	fourFours := Take(Repeat(4), 4)
	printNextN(fourFours, 5)

	// Output:
	// 4 true
	// 4 true
	// 4 true
	// 4 true
	// 0 false
}

func ExampleRepeatWith() {
	curr := 1
	pow2 := Take(RepeatWith(func() int {
		tmp := curr
		curr *= 2
		return tmp
	}), 4)
	printNextN(pow2, 5)

	// Output:
	// 1 true
	// 2 true
	// 4 true
	// 8 true
	// 0 false
}

func ExampleSuccessors() {
	powersOf10 := Successors(1, func(n uint) (uint, bool) {
		_, lo := bits.Mul(n, 10)
		if lo >= math.MaxUint16 {
			return 0, false
		}

		return lo, true
	})
	fmt.Println(Collect(powersOf10))

	// Output:
	// [1 10 100 1000 10000]
}

func ExampleSum() {
	fmt.Println(Sum(New(1, 2, 3)))
	// Output:
	// 6
}

func ExampleTake() {
	iter := Take(New(1, 2, 3), 2)
	printNextN(iter, 3)

	// Output:
	// 1 true
	// 2 true
	// 0 false
}

func ExampleZip() {
	iter := Zip(New(1, 2, 3), New(4, 5, 6))
	printNextN(iter, 4)

	// Output:
	// {1 4} true
	// {2 5} true
	// {3 6} true
	// {0 0} false
}
