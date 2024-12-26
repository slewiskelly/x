# conc

Package `conc` implements equivalents to the structured and unstructued
concurrency primitives offered by the Verse programming language.

## Usage

```go
import "github.com/slewiskelly/x/conc"
```

  ```go
// Branch
//
// Functions a, b, and c are all immediately executed and Branch immediately
// returns without waiting for any of them to complete.
conc.Branch(ctx, a, b, c)

// Race
//
// Functions a, b, and c are all immediately executed and Race blocks until
// the first completes, returning its result and sending a cancellation signals
// to the others.
got, err := conc.Race(ctx, a, b, c)

// Rush
//
// Functions a, b, and c are all immediately executed and Race blocks until
// the first completes, returning its result. Unlike Race a cancellation signal
// is not sent to the others.
got, err := conc.Rush(ctx, a, b, c)

// Spawn
//
// Function a is immediately executed and Spawn returns a corresponding
// Task, without waiting for it to complete.
task := conc.Spawn(ctx, a)

// Use the returned Task's Await method to retrieve the result.
got, err := task.Await(context.WithTimeout(ctx, 30*time.Second))

// Sync
//
// Functions a, b, and c are all immediately executed and Race blocks until
// the first completes, returning its result. Unlike Race a cancellation signal
// is not sent to the others.
got, err := conc.Sync(ctx, a, b, c)
```
