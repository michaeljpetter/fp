package fp

func Project2[T1, T2, U1, U2 any](project1 func(T1) U1, project2 func(T2) U2) func(T1, T2) (U1, U2) {
	return func(t1 T1, t2 T2) (U1, U2) {
		return project1(t1), project2(t2)
	}
}
