# set

Package `set` implements a container that holds a set of elements.

## Usage

```go
import "github.com/slewiskelly/x/container/set"
```

```go
// Set
//
// Implements a container which holds a set of elements.
//
// Reading from a set is safe for concurrent use. However, modification of a set
// is not.
a, b := set.New(1, 2, 3, 4, 5, 6), set.New(4, 5, 6, 7, 8, 9)

// Difference
//
// Returns a new set containing all values in a that are not present in b.
difference := set.Difference(a, b)

// Intersection
//
// Returns a new set containing the intersection of a and b.
intersection := set.Intersection(a, b)

// Union
//
// Union returns a new set containing the union of a and b.
union := set.Union(a, b)
```
