package fp

// Always returns a unary function that always returns the given value.
func Always[T, U any](u U) func(T) U {
	return func(T) U {
		return u
	}
}

// Always2 returns a binary function that always returns the given value.
func Always2[T1, T2, U any](u U) func(T1, T2) U {
	return func(T1, T2) U {
		return u
	}
}

// AlwaysTo2 returns a unary function that always returns the two given values.
func AlwaysTo2[T, U1, U2 any](u1 U1, u2 U2) func(T) (U1, U2) {
	return func(T) (U1, U2) {
		return u1, u2
	}
}

// Always2To2 returns a binary function that always returns the two given values.
func Always2To2[T1, T2, U1, U2 any](u1 U1, u2 U2) func(T1, T2) (U1, U2) {
	return func(T1, T2) (U1, U2) {
		return u1, u2
	}
}
