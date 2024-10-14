// Package types defines a set of constraint types
// used with type parameters in the fp package.
package types

// SignedInt is a constraint that includes all signed integer types.
type SignedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// UnsignedInt is a constraint that includes all unsigned integer types.
type UnsignedInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Int is a constraint that includes all signed and unsigned integer types.
type Int interface {
	SignedInt | UnsignedInt
}

// Float is a constraint that includes all floating-point types.
type Float interface {
	~float32 | ~float64
}

// Number is a constraint that includes all integer and floating-point types.
type Number interface {
	Int | Float
}

// Ordered is a constraint that includes all ordered types:
// those that support the operators < <= >= >.
type Ordered interface {
	Number | ~string
}
