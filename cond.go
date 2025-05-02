package fp

// Cond encapsulates conditional (if/else) logic, expressed by a series of unary predicate/result pairs.
// The [Cond.Apply] method is used to execute the conditional logic against a value.
type Cond[T, U any] []struct {
	Predicate func(T) bool
	Result    func(T) U
}

// Apply applies a value to each predicate in order until a true is yielded,
// then applies the same value to the result function paired with that predicate
// to produce the final value. If no predicate returns true, a zero value is returned.
func (c Cond[T, U]) Apply(t T) (zero U) {
	for _, pair := range c {
		if pair.Predicate(t) {
			return pair.Result(t)
		}
	}
	return
}

// Cond encapsulates conditional (if/else) logic, expressed by a series of binary predicate/result pairs.
// The [Cond2.Apply] method is used to execute the conditional logic against two values.
type Cond2[T1, T2, U any] []struct {
	Predicate func(T1, T2) bool
	Result    func(T1, T2) U
}

// Apply applies two values to each predicate in order until a true is yielded,
// then applies the same two values to the result function paired with that predicate
// to produce the final value. If no predicate returns true, a zero value is returned.
func (c Cond2[T1, T2, U]) Apply(t1 T1, t2 T2) (zero U) {
	for _, pair := range c {
		if pair.Predicate(t1, t2) {
			return pair.Result(t1, t2)
		}
	}
	return
}

// CondTo2 encapsulates conditional (if/else) logic, expressed by a series of unary predicate/1-to-2 result pairs.
// The [CondTo2.Apply] method is used to execute the conditional logic against a value.
type CondTo2[T, U1, U2 any] []struct {
	Predicate func(T) bool
	Result    func(T) (U1, U2)
}

// Apply applies a value to each predicate in order until a true is yielded,
// then applies the same value to the result function paired with that predicate
// to produce the final values. If no predicate returns true, zero values are returned.
func (c CondTo2[T, U1, U2]) Apply(t T) (zero1 U1, zero2 U2) {
	for _, pair := range c {
		if pair.Predicate(t) {
			return pair.Result(t)
		}
	}
	return
}

// Cond2To2 encapsulates conditional (if/else) logic, expressed by a series of binary predicate/2-to-2 result pairs.
// The [Cond2To2.Apply] method is used to execute the conditional logic against two values.
type Cond2To2[T1, T2, U1, U2 any] []struct {
	Predicate func(T1, T2) bool
	Result    func(T1, T2) (U1, U2)
}

// Apply applies two values to each predicate in order until a true is yielded,
// then applies the same two values to the result function paired with that predicate
// to produce the final values. If no predicate returns true, zero values are returned.
func (c Cond2To2[T1, T2, U1, U2]) Apply(t1 T1, t2 T2) (zero1 U1, zero2 U2) {
	for _, pair := range c {
		if pair.Predicate(t1, t2) {
			return pair.Result(t1, t2)
		}
	}
	return
}
