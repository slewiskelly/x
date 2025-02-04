// Package slices implements various functions for slices of any type and are
// omitted from the standard library.
package slices

// Every returns true if the given function returns true for all elements in the
// given slice; otherwise false.
func Every[S ~[]T, T any](s S, fn func(e T) bool) bool {
	for _, e := range s {
		if !fn(e) {
			return false
		}
	}

	return true
}

// Filter returns a slice populated with all elements in which the given
// function returns true for.
func Filter[S ~[]T, T any](s S, fn func(e T) bool) S {
	var x S

	for _, e := range s {
		if fn(e) {
			x = append(x, e)
		}
	}

	return x
}

// Find returns the first element and true for which the given function returns
// true for; otherwise the type's zero value and false.
func Find[S ~[]T, T any](s S, fn func(e T) bool) (T, bool) {
	if i := FindIndex(s, fn); i != -1 {
		return s[i], true
	}

	var x T
	return x, false
}

// FindIndex returns the index of the first element for which the given function
// returns true for; otherwise -1.
func FindIndex[S ~[]T, T any](s S, fn func(e T) bool) int {
	for i, e := range s {
		if fn(e) {
			return i
		}
	}

	return -1
}

// ForEach executes the given function for each element in the given slice.
func ForEach[S ~[]T, T any](s S, fn func(e T)) {
	for _, e := range s {
		fn(e)
	}
}

// Map returns a slice populated with the result of executing the given function
// for all elements of the given slice.
func Map[S ~[]T, T any](s S, fn func(e T) T) S {
	x := make(S, len(s))

	for i, e := range s {
		x[i] = fn(e)
	}

	return x
}

// Reduce returns the result of executing the given function for all elements of
// the given slice.
//
// Each execution is passed the result of the previous execution, with the
// exception of the first, which is passed the given initial value.
func Reduce[S ~[]T, T any](s S, fn func(a, e T) T, initial T) T {
	x := initial

	for _, e := range s {
		x = fn(x, e)
	}

	return x
}

// Some returns true if the given function returns true for any of the elements
// in the given slice; otherwise false.
func Some[S ~[]T, T any](s S, fn func(e T) bool) bool {
	for _, e := range s {
		if fn(e) {
			return true
		}
	}

	return false
}
