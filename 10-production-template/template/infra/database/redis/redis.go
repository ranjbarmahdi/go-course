package redis

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"template/infra/config"

	goredis "github.com/redis/go-redis/v9"
)

func newClient(cfg *config.Config) (*goredis.Client, error) {
	if cfg.RedisSelfHost() {
		opt, err := goredis.ParseURL(cfg.Redis.URL)
		if err != nil {
			return nil, fmt.Errorf("parse redis url: %w", err)
		}
		return goredis.NewClient(opt), nil
	}

	return goredis.NewFailoverClient(&goredis.FailoverOptions{
		MasterName: cfg.Redis.SentinelMaster,
		SentinelAddrs: []string{
			fmt.Sprintf("%s:%d", cfg.Redis.SentinelHost, cfg.Redis.SentinelPort),
		},
		Username:         cfg.Redis.SentinelUser,
		Password:         cfg.Redis.SentinelPass,
		SentinelPassword: cfg.Redis.SentinelPass,
	}), nil
}

func ping(ctx context.Context, client *goredis.Client) error {
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return client.Ping(pingCtx).Err()
}

func closeClient(client *goredis.Client) error {
	if client == nil {
		return nil
	}
	return client.Close()
}

// NewRedis is the Wire provider — same shape as postgres.NewPostgres.
var log = slog.Default().With("component", "redis")

func NewRedis(ctx context.Context, cfg *config.Config) (*goredis.Client, func(), error) {
	mode := "sentinel"
	if cfg.RedisSelfHost() {
		mode = "direct"
	}
	client, err := newClient(cfg)
	if err != nil {
		return nil, nil, err
	}
	if err := ping(ctx, client); err != nil {
		_ = closeClient(client)
		return nil, nil, fmt.Errorf("ping redis: %w", err)
	}
	log.Info("connected", "mode", mode)
	cleanup := func() {
		if err := closeClient(client); err != nil {
			log.Warn("disconnect failed", "err", err)
			return
		}
		log.Info("disconnected")
	}
	return client, cleanup, nil
}
