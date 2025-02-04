# slices

Package `slices` implements various functions for slices of any type and are
omitted from the standard library.

## Usage

```go
import "github.com/slewiskelly/x/slices"
```

```go
// Every
//
// Returns true if the given function returns true for all elements in the
// given slice; otherwise false.
ok := Every([]int{1, 30, 39, 29, 10, 13}, func(e int) bool {
	return e < 40
})

// Filter
//
// Returns a slice populated with all elements in which the given
// function returns true for.
x := Filter([]string{"spray", "elite", "exuberant", "destruction", "present"}, func(e string) bool {
	return len(e) > 6
})

// Find
//
// Returns the first element and true for which the given function returns
// true for; otherwise the type's zero value and false.
x, ok := Find([]int{5, 12, 8, 130, 44}, func(e int) bool {
	return e > 10
})

// FindIndex
//
// Returns the index of the first element for which the given function
// returns true for; otherwise -1.
i := FindIndex([]int{5, 12, 8, 130, 44}, func(e int) bool {
	return e > 13
})

// ForEach
//
// Executes the given function for each element in the given slice.
ForEach([]string{"a", "b", "c"}, func(e string) {
	fmt.Println(strings.ToUpper(e))
})

// Map
//
// Returns a slice populated with the result of executing the given function
// for all elements of the given slice.
x := Map([]int{1, 4, 9, 16}, func(e int) int {
	return e * 2
})

// Reduce
//
// Returns the result of executing the given function for all elements of
// the given slice.
//
// Each execution is passed the result of the previous execution, with the
// exception of the first, which is passed the given initial value.
sum := Reduce([]int{1, 2, 3, 4}, func(a int, e int) int {
	return a + e
}, 0)

// Some
//
// Returns true if the given function returns true for any of the elements
// in the given slice; otherwise false.
ok := Some([]int{1, 2, 3, 4, 5}, func(e int) bool {
	return e%2 == 0
})
```
