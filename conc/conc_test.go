package conc

import (
	"context"
	"slices"
	"sync/atomic"
	"testing"
	"time"
)

func TestBranch(t *testing.T) {
	var count atomic.Int32

	fn := func(ctx context.Context) {
		count.Add(1)
	}

	Branch(context.Background(), fn, fn, fn)

	<-time.After(3 * time.Second)

	if got, want := count.Load(), int32(3); got != want {
		t.Errorf("Branch(ctx, _, _, _) incremented counter %d times; want %d", got, want)
	}
}

func TestRace(t *testing.T) {
	got, err := Race(context.Background(),
		func(ctx context.Context) (string, error) {
			<-time.After(3 * time.Second)
			return "slow", nil
		},
		func(ctx context.Context) (string, error) {
			<-time.After(2 * time.Second)
			return "mid", nil
		},
		func(ctx context.Context) (string, error) {
			<-time.After(1 * time.Second)
			return "fast", nil
		},
	)
	if err != nil {
		t.Errorf("Race(ctx, ...) = _, %v", err)
	}

	if want := "fast"; got != want {
		t.Errorf("Race(ctx, ...) = %q, _; want %q", got, want)
	}
}

func TestRush(t *testing.T) {
	got, err := Race(context.Background(),
		func(ctx context.Context) (string, error) {
			<-time.After(3 * time.Second)
			return "slow", nil
		},
		func(ctx context.Context) (string, error) {
			<-time.After(2 * time.Second)
			return "mid", nil
		},
		func(ctx context.Context) (string, error) {
			<-time.After(1 * time.Second)
			return "fast", nil
		},
	)
	if err != nil {
		t.Errorf("Rush(ctx, ...) = _, %v", err)
	}

	if want := "fast"; got != want {
		t.Errorf("Rush(ctx, ...) = %q, _; want %q", got, want)
	}
}

func TestSpawn(t *testing.T) {
	task := Spawn(context.Background(), func(ctx context.Context) (string, error) {
		return "hello, world", nil
	})

	got, err := task.Await(context.Background())
	if err != nil {
		t.Errorf("Spawn(ctx, ...).Await(ctx) = _, %v", err)
	}

	if want := "hello, world"; got != want {
		t.Errorf("Spawn(ctx, ...).Await(ctx) = %q, _; want %q", got, want)
	}
}

func TestSync(t *testing.T) {
	got, err := Sync(context.Background(),
		func(ctx context.Context) (int, error) {
			return 1, nil
		},
		func(ctx context.Context) (int, error) {
			return 2, nil
		},
		func(ctx context.Context) (int, error) {
			return 3, nil
		},
	)
	if err != nil {
		t.Errorf("Sync(ctx, ...) = _, %v", err)
	}

	slices.Sort(got)

	want := []int{1, 2, 3}

	if !slices.Equal(got, want) {
		t.Errorf("Sync(ctx, ...) = %v, _; want %v", got, want)
	}
}
