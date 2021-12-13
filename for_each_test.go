package iter_test

import (
	"fmt"

	. "github.com/fdschonborn/go-iter"
)

func ExampleForEach() {
	a := []int{1, 2, 3}
	ForEach(FromSlice(a), func(x int) {
		fmt.Println(x * 2)
	})

	// Output:
	// 2
	// 4
	// 6
}
