package fp

// Tap returns a unary function that executes a given interceptor function
// with the input value it receives, then returns the input value.
func Tap[T any](interceptor func(T)) func(T) T {
	return func(t T) T {
		interceptor(t)
		return t
	}
}

// Tap2 returns a binary function that executes a given interceptor function
// with the input values it receives, then returns the input values.
func Tap2[T1, T2 any](interceptor func(T1, T2)) func(T1, T2) (T1, T2) {
	return func(t1 T1, t2 T2) (T1, T2) {
		interceptor(t1, t2)
		return t1, t2
	}
}
