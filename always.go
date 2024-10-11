package fp

func Always[T any](t T) func() T {
	return func() T {
		return t
	}
}

func Always2[T1, T2 any](t1 T1, t2 T2) func() (T1, T2) {
	return func() (T1, T2) {
		return t1, t2
	}
}
