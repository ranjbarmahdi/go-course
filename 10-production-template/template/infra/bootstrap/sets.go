package bootstrap

import (
	"template/infra/database/postgres"
	"template/infra/database/redis"
	"template/infra/httpserver"
	"template/infra/runtime"

	"github.com/google/wire"
)

var AdapterSet = wire.NewSet(
	postgres.NewRunInTx,
	provideIDGenerator,
)

var DatabaseSet = wire.NewSet(
	postgres.NewPostgres,
	redis.NewRedis,
)

var IndicatorsSet = wire.NewSet(
	postgres.NewIndicator,
	redis.NewIndicator,
)

var HttpSet = wire.NewSet(
	provideIndicators,
	provideAddr,
	httpserver.NewRouter,
	httpserver.NewServer,
)

var RuntimeSet = wire.NewSet(
	provideRunners,
	newApp,
)

// newApp wraps variadic NewApp — Wire handles []Runner better this way
func newApp(runners []runtime.Runner) *runtime.App {
	return runtime.NewApp(runners...)
}
