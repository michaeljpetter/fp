package fp

// Project2To2 creates a 2-to-2 projection out of two 1-to-1 projections.
func Project2To2[T1, T2, U1, U2 any](project1 func(T1) U1, project2 func(T2) U2) func(T1, T2) (U1, U2) {
	return func(t1 T1, t2 T2) (U1, U2) {
		return project1(t1), project2(t2)
	}
}
