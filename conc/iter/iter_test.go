package iter

import (
	"context"
	"fmt"
)

func ExampleForEach() {
	_ = ForEach(context.Background(), []string{"a", "b", "c"}, func(ctx context.Context, i int, e string) error {
		fmt.Printf("%d: %s\n", i, e)

		return nil
	}, Limit(1))
	// Output: 0: a
	// 1: b
	// 2: c
}
