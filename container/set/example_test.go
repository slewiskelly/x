package set_test

import (
	"fmt"
	"slices"

	"github.com/slewiskelly/x/container/set"
)

func ExampleDifference() {
	a, b := set.New(1, 2, 3, 4, 5, 6), set.New(4, 5, 6, 7, 8, 9)

	difference := set.Difference(a, b)

	fmt.Println(slices.Sorted(difference.Values()))
	// Output: [1 2 3]
}

func ExampleIntersection() {
	a, b := set.New(1, 2, 3, 4, 5, 6), set.New(4, 5, 6, 7, 8, 9)

	intersection := set.Intersection(a, b)

	fmt.Println(slices.Sorted(intersection.Values()))
	// Output: [4 5 6]
}

func ExampleUnion() {
	a, b := set.New(1, 2, 3, 4, 5, 6), set.New(4, 5, 6, 7, 8, 9)

	union := set.Union(a, b)

	fmt.Println(slices.Sorted(union.Values()))
	// Output: [1 2 3 4 5 6 7 8 9]
}
