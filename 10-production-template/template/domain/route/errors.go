package route

import "template/domain/domainerror"

var (
	ErrNotFound              = domainerror.New(domainerror.NotFound, "route not found")
	ErrInvalidName           = domainerror.New(domainerror.InvalidInput, "invalid route name")
	ErrInvalidVesselID       = domainerror.New(domainerror.InvalidInput, "invalid vessel id")
	ErrInvalidCreatorID      = domainerror.New(domainerror.InvalidInput, "invalid creator id")
	ErrInvalidCreatorType    = domainerror.New(domainerror.InvalidInput, "invalid creator type")
	ErrInvalidCreatedAt      = domainerror.New(domainerror.InvalidInput, "created at is required")
	ErrInvalidFileType       = domainerror.New(domainerror.InvalidInput, "invalid file type")
	ErrInvalidCoordinates    = domainerror.New(domainerror.InvalidInput, "invalid coordinates")
	ErrInvalidOrderNumber    = domainerror.New(domainerror.InvalidInput, "invalid order number")
	ErrInvalidWaypointType   = domainerror.New(domainerror.InvalidInput, "invalid waypoint type")
	ErrDuplicateOrderNumber  = domainerror.New(domainerror.Conflict, "duplicate order number")
	ErrWaypointRouteMismatch = domainerror.New(domainerror.Domain, "waypoint route id mismatch")
	ErrInvalidRouteID        = domainerror.New(domainerror.InvalidInput, "route id is required")
	ErrInvalidWaypointID     = domainerror.New(domainerror.InvalidInput, "waypoint id is required")
)
