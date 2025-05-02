package fp

// Flip2 wraps a binary function, reversing the order of its arguments.
func Flip2[T1, T2, U any](function func(T1, T2) U) func(T2, T1) U {
	return func(t2 T2, t1 T1) U {
		return function(t1, t2)
	}
}

// Flip2To2 wraps a binary function, reversing the order of its arguments.
func Flip2To2[T1, T2, U1, U2 any](function func(T1, T2) (U1, U2)) func(T2, T1) (U1, U2) {
	return func(t2 T2, t1 T1) (U1, U2) {
		return function(t1, t2)
	}
}
