package fp

// Not wraps a unary predicate, inverting its result.
func Not[T any](predicate func(T) bool) func(T) bool {
	return func(t T) bool {
		return !predicate(t)
	}
}

// Not2 wraps a binary predicate, inverting its result.
func Not2[T1, T2 any](predicate func(T1, T2) bool) func(T1, T2) bool {
	return func(t1 T1, t2 T2) bool {
		return !predicate(t1, t2)
	}
}
