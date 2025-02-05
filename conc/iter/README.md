# iter

Package `iter` implements concurrent iterators.

## Usage

```go
import "github.com/slewiskelly/x/conc/iter"
```

```go
// ForEach
//
// Executes the given function for each element in the given slice.
//
// Execution upon each element may occur concurrently, depending on the limit.
err := ForEach(context.Background(), []string{"a", "b", "c"}, func(ctx context.Context, i int, e string) error {
	fmt.Printf("%d: %s\n", i, e)

	return nil
})
```
