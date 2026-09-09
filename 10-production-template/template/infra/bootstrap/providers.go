package bootstrap

import (
	"fmt"

	"template/application/contracts"
	"template/infra/adapters"
	"template/infra/config"
	"template/infra/database/postgres"
	"template/infra/database/redis"
	"template/infra/httpserver"
	"template/infra/runtime"
)

func provideIndicators(postgres *postgres.Indicator, redis *redis.Indicator) []runtime.Indicator {
	return []runtime.Indicator{postgres, redis}
}

func provideAddr(cfg *config.Config) httpserver.Addr {
	return httpserver.Addr(fmt.Sprintf(":%d", cfg.Port))
}

func provideRunners(server *httpserver.Server) []runtime.Runner {
	return []runtime.Runner{server}
}

func provideIDGenerator() contracts.IDGenerator {
	return adapters.UUIDGenerator{}
}
