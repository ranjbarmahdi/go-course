//go:build wireinject
// +build wireinject

package bootstrap

import (
	"context"
	"errors"

	"template/infra/config"
	"template/infra/runtime"

	"github.com/google/wire"
)

var ErrWireNotGenerated = errors.New("wire dependencies are not generated; run wire in infra/bootstrap")

func Wire(ctx context.Context, cfg *config.Config) (*runtime.App, func(), error) {
	wire.Build(
		DatabaseSet,
		IndicatorsSet,
		HttpSet,
		RuntimeSet,
		// AdapterSet,
	)
	return nil, nil, ErrWireNotGenerated
}
