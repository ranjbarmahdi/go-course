package config

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type RedisConfig struct {
	URL            string `env:"REDIS_URL"`
	SentinelHost   string `env:"REDIS_SENTINEL_HOST"`
	SentinelPort   int    `env:"REDIS_SENTINEL_PORT" envDefault:"26379"`
	SentinelUser   string `env:"REDIS_SENTINEL_USER"`
	SentinelMaster string `env:"REDIS_SENTINEL_MASTER"`
	SentinelPass   string `env:"REDIS_SENTINEL_PASS"`
}

type Config struct {
	Port        int    `env:"PORT" envDefault:"8080"`
	Env         string `env:"ENV" envDefault:"development"`
	DatabaseURL string `env:"DATABASE_URL,required"`
	DBSchema    string `env:"DB_SCHEMA,required"`
	LogLevel    string `env:"LOG_LEVEL" envDefault:"info"`
	AutoMigrate bool   `env:"AUTO_MIGRATE" envDefault:"false"`
	Redis       RedisConfig
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("validate config: %w", err)
	}

	return cfg, nil
}

func (c *Config) RedisSelfHost() bool {
	return c.Env == "self-host" || c.Env == "development"
}

func (r RedisConfig) Validate(selfHost bool) error {
	if selfHost {
		if strings.TrimSpace(r.URL) == "" {
			return fmt.Errorf("REDIS_URL is required when ENV is self-host or development")
		}
		return nil
	}
	if strings.TrimSpace(r.SentinelHost) == "" {
		return fmt.Errorf("REDIS_SENTINEL_HOST is required in sentinel mode")
	}
	if r.SentinelPort < 1 || r.SentinelPort > 65535 {
		return fmt.Errorf("REDIS_SENTINEL_PORT must be between 1 and 65535")
	}
	if strings.TrimSpace(r.SentinelMaster) == "" {
		return fmt.Errorf("REDIS_SENTINEL_MASTER is required in sentinel mode")
	}
	if strings.TrimSpace(r.SentinelPass) == "" {
		return fmt.Errorf("REDIS_SENTINEL_PASS is required in sentinel mode")
	}
	return nil
}

func (c *Config) Validate() error {
	if c.Port < 1 || c.Port > 65535 {
		return fmt.Errorf("PORT must be between 1 and 65535")
	}

	if strings.TrimSpace(c.DBSchema) == "" {
		return fmt.Errorf("DB_SCHEMA is required")
	}

	if err := c.Redis.Validate(c.RedisSelfHost()); err != nil {
		return err
	}

	return nil
}

func ParseSlogLevel(level string) slog.Level {
	switch strings.ToLower(strings.TrimSpace(level)) {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
