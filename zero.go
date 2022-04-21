package iter

// Zero returns the zero value of type T.
//
// This is a equivalent to doing:
//
//     var zero T
//     return zero
func Zero[T any]() T {
	var zero T
	return zero
}
