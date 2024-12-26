// Package conc implements equivalents to the structured and unstructued
// concurrency primitives offered by the Verse programming language.
//
// WARNING: This package is not suitable for production use.
package conc

import (
	"context"
	"errors"
	"sync"
)

// Task represents the execution of a function.
type Task[T any] struct {
	ch chan result[T]
}

type result[T any] struct {
	x T
	e error
}

// Await blocks until a result is made available and returns it.
//
// Cancelling the given context, before a result is made available, causes Await
// to return the context's error.
func (t *Task[T]) Await(ctx context.Context) (T, error) {
	select {
	case <-ctx.Done():
		var x T
		return x, ctx.Err()
	case resp := <-t.ch:
		return resp.x, resp.e
	}
}

// Branch immediately executes all given functions concurrently, it does not
// block.
func Branch(ctx context.Context, fn ...func(ctx context.Context)) {
	for _, f := range fn {
		go f(ctx)
	}
}

// Race immediately executes all given functions concurrently, it blocks until
// a result is made available.
//
// Race returns the first result and sends a cancellation signal to all other
// executing functions.
//
// Cancelling the given context, before a result is made available, causes Race
// to return the context's error.
func Race[T any](ctx context.Context, fn ...func(ctx context.Context) (T, error)) (T, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	return Rush(ctx, fn...)
}

// Rush immediately executes all given functions concurrently, it blocks until
// a result is made available.
//
// Rush returns the first result and, unlike Race, a cancellation signal is not
// sent to other executing functions.
//
// Cancelling the given context, before a result is made available, causes Rush
// to return the context's error.
func Rush[T any](ctx context.Context, fn ...func(ctx context.Context) (T, error)) (T, error) {
	ch := make(chan result[T], len(fn))

	for _, f := range fn {
		go func() {
			x, err := f(ctx)
			ch <- result[T]{x, err}
		}()
	}

	select {
	case <-ctx.Done():
		var x T
		return x, ctx.Err()
	case resp := <-ch:
		return resp.x, resp.e
	}
}

// Spawn immediately executes the given function, it does not block and instead
// returns a Task.
//
// To retrieve the result, call Await on the returned Task, which will block
// until a result is made available.
func Spawn[T any](ctx context.Context, fn func(ctx context.Context) (T, error)) *Task[T] {
	t := &Task[T]{make(chan result[T], 1)}

	go func() {
		x, err := fn(ctx)

		t.ch <- result[T]{x, err}
		close(t.ch)
	}()

	return t
}

// Sync immediately executes all given functions concurrently, it blocks until
// they have all completed.
//
// Sync returns all results and a single error, the union of all errors
// encountered.
func Sync[T any](ctx context.Context, fn ...func(ctx context.Context) (T, error)) ([]T, error) {
	var resp []T
	var errs []error

	ch := make(chan result[T], len(fn))

	var wg sync.WaitGroup

	for _, f := range fn {
		wg.Add(1)

		go func() {
			defer wg.Done()

			x, err := f(ctx)
			ch <- result[T]{x, err}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for r := range ch {
		resp = append(resp, r.x)
		errs = append(errs, r.e)
	}

	return resp, errors.Join(errs...)
}
