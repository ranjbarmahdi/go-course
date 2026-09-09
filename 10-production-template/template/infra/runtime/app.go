package runtime

import (
	"context"
	"log/slog"

	"golang.org/x/sync/errgroup"
)

type App struct {
	runners []Runner
}

func NewApp(runners ...Runner) *App {
	return &App{
		runners: runners,
	}
}

func (a *App) Run(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)

	for _, runner := range a.runners {
		slog.Info("starting runner", "name", runner.Name())
		group.Go(func() error {
			return runner.Start(ctx)
		})
	}

	return group.Wait()
}
