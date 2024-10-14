package fp

import (
	"reflect"
	"sync"
)

// Once wraps any function, ensuring it is executed only once,
// and returning its cached results (including panics) on every call thereafter.
// The returned function may be called concurrently.
func Once[F any](function F) F {
	var once sync.Once
	var out []reflect.Value
	var p any

	return reflect.MakeFunc(
		reflect.TypeFor[F](),
		func(args []reflect.Value) []reflect.Value {
			once.Do(func() {
				defer func() { p = recover() }()
				out = reflect.ValueOf(function).Call(args)
			})
			if p != nil {
				panic(p)
			}
			return out
		},
	).Interface().(F)
}
