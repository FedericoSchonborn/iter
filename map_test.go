package iter_test

import (
	"fmt"

	. "github.com/fdschonborn/go-iter"
)

func ExampleMap() {
	a := []int{1, 2, 3}
	iter := Map(FromSlice(a), func(x int) int {
		return 2 * x
	})

	fmt.Println(iter.Next())
	fmt.Println(iter.Next())
	fmt.Println(iter.Next())
	// Output:
	// 2 true
	// 4 true
	// 6 true
}
