package contracts

import "context"

type RunInTx[T any] func(ctx context.Context, fn func(ctx context.Context) T) T
