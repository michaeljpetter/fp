package fp

// Always returns a function that always returns the given value.
func Always[T any](t T) func() T {
	return func() T {
		return t
	}
}

// Always2 returns a function that always returns the given values.
func Always2[T1, T2 any](t1 T1, t2 T2) func() (T1, T2) {
	return func() (T1, T2) {
		return t1, t2
	}
}
