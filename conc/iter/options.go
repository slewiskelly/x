package iter

// Option is an option to iterators.
type Option interface {
	apply(*options)
}

// Limit sets the parallelization limit, by default there is no limit.
//
// A negative value also disables any limits.
func Limit(i int) Option {
	return option(func(o *options) {
		o.limit = i
	})
}

type options struct {
	limit int
}

type option func(*options)

func (o option) apply(opts *options) {
	o(opts)
}
