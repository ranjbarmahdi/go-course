package runtime

import "context"

type Runner interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Name() string
}

type Indicator interface {
	Name() string
	Ready(ctx context.Context) bool
}
