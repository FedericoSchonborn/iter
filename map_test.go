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

	for i := 0; i < 3; i++ {
		j, _ := iter.Next()
		fmt.Println(j)
	}
	// Output:
	// 2
	// 4
	// 6
}
