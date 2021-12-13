package iter_test

import (
	"fmt"

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

	iter := From(1, 2, 3, 4)

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
	fmt.Println(All(From(1, 2, 3), greaterThan(0)))
	fmt.Println(All(From(1, 2, 3), greaterThan(2)))
	// Output:
	// true
	// false
}

func ExampleAny() {
	fmt.Println(Any(From(1, 2, 3), greaterThan(0)))
	fmt.Println(Any(From(1, 2, 3), greaterThan(5)))
	// Output:
	// true
	// false
}

func ExampleCount() {
	fmt.Println(Count(From(1, 2, 3)))
	fmt.Println(Count(From(1, 2, 3, 4, 5)))
	// Output:
	// 3
	// 5
}

func ExampleEnumerate() {
	iter := Enumerate(From('a', 'b', 'c'))
	printNextN(iter, 4)

	// Output:
	// {0 97} true
	// {1 98} true
	// {2 99} true
	// {0 0} false
}

func ExampleFilter() {
	iter := Filter(From(0, 1, 2), isPositive)
	printNextN(iter, 3)

	// Output:
	// 1 true
	// 2 true
	// 0 false
}

func ExampleFold() {
	sum := Fold(From(1, 2, 3), 0, func(acc int, x int) int {
		return acc + x
	})

	fmt.Println(sum)
	// Output:
	// 6
}

func ExampleForEach() {
	ForEach(From(1, 2, 3), func(x int) {
		fmt.Println(x * 2)
	})

	// Output:
	// 2
	// 4
	// 6
}

func ExampleMap() {
	iter := Map(From(1, 2, 3), double)
	printNextN(iter, 3)

	// Output:
	// 2 true
	// 4 true
	// 6 true
}

func ExampleIntoSlice() {
	fmt.Println(IntoSlice(Map(From(1, 2, 3), double)))
	// Output:
	// [2 4 6]
}
