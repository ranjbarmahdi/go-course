package bootstrap

import (
	createsample "template/application/usecase/sample/create-sample"
	getsample "template/application/usecase/sample/get-sample"
	"template/infra/adapters/repository"
	"template/infra/database/postgres"
	"template/infra/database/redis"
	"template/infra/httpserver"
	"template/infra/runtime"

	domainsample "template/domain/sample"
	samplehttp "template/infra/httpserver/routes/sample"

	"github.com/google/wire"
)

var AdapterSet = wire.NewSet(
	postgres.NewRunInTx,
	provideIDGenerator,
	repository.NewSampleRepository,
	wire.Bind(new(domainsample.SampleRepository), new(*repository.SampleRepository)),
)

var DatabaseSet = wire.NewSet(
	postgres.NewPostgres,
	redis.NewRedis,
)

var IndicatorsSet = wire.NewSet(
	postgres.NewIndicator,
	redis.NewIndicator,
)

var SampleUseCaseSet = wire.NewSet(
	createsample.New,
	wire.Bind(new(createsample.UseCase), new(*createsample.Implementation)),
	getsample.New,
	wire.Bind(new(getsample.UseCase), new(*getsample.Implementation)),
)

var SampleHandlerSet = wire.NewSet(
	samplehttp.NewHandler,
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
