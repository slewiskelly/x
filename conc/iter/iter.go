// Package iter implements concurrent iterators.
package iter

import (
	"context"

	"golang.org/x/sync/errgroup"
)

// ForEach executes the given function for each element in the given slice.
//
// Execution upon each element may occur concurrently, depending on the limit.
//
// For each iteration the given function is passed:
// - A child context derived from the context passed to ForEach.
// - The index of the element
// - The element itself
//
// ForEach blocks until either all iterations have completed or an error is
// encountered; with the first error being encountered being returned.
//
// If the parent context is canceled while there are still iterations remaining,
// ForEach will immediately return with the cancellation cause.
func ForEach[S ~[]E, E any](ctx context.Context, s S, fn func(ctx context.Context, i int, e E) error, opts ...Option) error {
	o := &options{
		limit: -1, // No limits by default.
	}

	for _, opt := range opts {
		opt.apply(o)
	}

	grp, ctx := errgroup.WithContext(ctx)
	grp.SetLimit(o.limit)

	for i, e := range s {
		select {
		case <-ctx.Done():
			return context.Cause(ctx)
		default:
		}

		grp.Go(func() error { return fn(ctx, i, e) })
	}

	return grp.Wait()
}
