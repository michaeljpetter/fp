package fp

func Flip2[T1, T2, U any](function func(T1, T2) U) func(T2, T1) U {
	return func(t2 T2, t1 T1) U {
		return function(t1, t2)
	}
}
