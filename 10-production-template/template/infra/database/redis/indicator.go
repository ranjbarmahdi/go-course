package redis

import (
	"context"
	"time"

	"template/infra/runtime"

	goredis "github.com/redis/go-redis/v9"
)

type Indicator struct {
	client *goredis.Client
}

func NewIndicator(client *goredis.Client) *Indicator {
	return &Indicator{client: client}
}

func (i *Indicator) Name() string {
	return "redis"
}

func (i *Indicator) Ready(ctx context.Context) bool {
	if i.client == nil {
		return false
	}

	pingCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	return i.client.Ping(pingCtx).Err() == nil
}

var _ runtime.Indicator = (*Indicator)(nil)
