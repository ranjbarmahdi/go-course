package route

import (
	"context"
	valueobjects "template/domain/value-objects"
)

type RouteRepository interface {
	CreateRoute(ctx context.Context, route Route) error
	FindByID(ctx context.Context, id valueobjects.RouteID) (*Route, error)
}
