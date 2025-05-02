package fp

// Over reduces a series of unary functions into a single function,
// which collects the results of applying each function to the
// argument it receives.
func Over[T, U any](functions ...func(T) U) func(T) []U {
	return func(t T) []U {
		out := make([]U, len(functions))
		for i, f := range functions {
			out[i] = f(t)
		}
		return out
	}
}

// Over2 reduces a series of binary functions into a single function,
// which collects the results of applying each function to the
// arguments it receives.
func Over2[T1, T2, U any](functions ...func(T1, T2) U) func(T1, T2) []U {
	return func(t1 T1, t2 T2) []U {
		out := make([]U, len(functions))
		for i, f := range functions {
			out[i] = f(t1, t2)
		}
		return out
	}
}
