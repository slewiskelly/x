// Package set implements a container that holds a set of elements.
package set

import (
	"fmt"
	"iter"
)

// Difference returns a new set containing all elements in a that are not present
// in b.
func Difference[T comparable](a, b *Set[T]) *Set[T] {
	x := &Set[T]{make(map[T]struct{})}

	for v := range a.Values() {
		if !b.Contains(v) {
			x.Add(v)
		}
	}

	return x
}

// Intersection returns a new set containing the intersection of a and b.
func Intersection[T comparable](a, b *Set[T]) *Set[T] {
	x := &Set[T]{make(map[T]struct{})}

	for v := range a.Values() {
		if b.Contains(v) {
			x.Add(v)
		}
	}

	return x
}

// Union returns a new set containing the union of a and b.
func Union[T comparable](a, b *Set[T]) *Set[T] {
	x := &Set[T]{make(map[T]struct{})}

	x.AddSet(a)
	x.AddSet(b)

	return x
}

// Set implements a container which holds a set of elements.
//
// Reading from a set is safe for concurrent use. However, modification of a set
// is not.
type Set[T comparable] struct {
	m map[T]struct{}
}

// New returns a new set containing the given elements.
func New[T comparable](e ...T) *Set[T] {
	s := &Set[T]{make(map[T]struct{})}

	for _, v := range e {
		s.m[v] = struct{}{}
	}

	return s
}

// Add adds element(s) to the set.
func (s *Set[T]) Add(e ...T) {
	for _, v := range e {
		s.m[v] = struct{}{}
	}
}

// AddSet adds all elements of the given set to the set.
func (s *Set[T]) AddSet(o *Set[T]) {
	for v := range o.Values() {
		s.Add(v)
	}
}

// All returns an iterator over (index, value) pairs of the set.
//
// Elements are yielded in an indeterminate order.
func (s *Set[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		var i int

		for v := range s.m {
			if !yield(i, v) {
				return
			}

			i++
		}
	}
}

// Clear clears all elements from the set.
func (s *Set[T]) Clear() {
	s.m = make(map[T]struct{})
}

// Clone returns a shallow copy of the set.
func (s *Set[T]) Clone() Set[T] {
	c := Set[T]{make(map[T]struct{})}

	c.AddSet(s)

	return c
}

// Contains returns true if the given value is contained in the set; otherwise
// false.
func (s *Set[T]) Contains(v T) bool {
	_, ok := s.m[v]
	return ok
}

// ContainsAll returns true if the set contains all of the elements in given
// set; otherwise false.
func (s *Set[T]) ContainsAll(o *Set[T]) bool {
	for v := range o.Values() {
		if _, ok := s.m[v]; !ok {
			return false
		}
	}

	return true
}

// ContainsAny returns true if the set contains any of the elements in the given
// set; otherwise false.
func (s *Set[T]) ContainsAny(o *Set[T]) bool {
	for v := range o.Values() {
		if _, ok := s.m[v]; ok {
			return true
		}
	}

	return false
}

// Equal returns true if the set is equal to the given set.
//
// Sets are considered equal if they are of the same length and all elements are
// equal.
func (s *Set[T]) Equal(o *Set[T]) bool {
	if s.Len() != o.Len() {
		return false
	}

	return s.ContainsAll(o)
}

// Filter removes any elements from the set in which the given function returns
// false.
func (s *Set[T]) Filter(keep func(T) bool) {
	for v := range s.Values() {
		if !keep(v) {
			delete(s.m, v)
		}
	}
}

// ForEach executes the given function for each element of the set.
func (s *Set[T]) ForEach(fn func(T)) {
	for v := range s.Values() {
		fn(v)
	}
}

// Len returns the number of elements in the set.
func (s *Set[T]) Len() int {
	return len(s.m)
}

// Remove removes the given elements from the set.
func (s *Set[T]) Remove(e ...T) {
	for _, v := range e {
		delete(s.m, v)
	}
}

// RemoveSet removes the elements of the given set from the set.
func (s *Set[T]) RemoveSet(o *Set[T]) {
	for v := range o.Values() {
		delete(s.m, v)
	}
}

// String returns a string representation of the set.
func (s *Set[T]) String() string {
	x := make([]T, 0, s.Len())

	for v := range s.Values() {
		x = append(x, v)
	}

	return fmt.Sprint(x)
}

// Values returns an iterator over elements of the set.
//
// Elements are yielded in an indeterminate order.
func (s *Set[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range s.m {
			if !yield(v) {
				return
			}
		}
	}
}
