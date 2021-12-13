package iter_test

import (
	"fmt"

	. "github.com/fdschonborn/go-iter"
)

func ExampleFold() {
	a := []int{1, 2, 3}
	sum := Fold(FromSlice(a), 0, func(acc int, x int) int {
		return acc + x
	})

	fmt.Println(sum)
	// Output:
	// 6
}
