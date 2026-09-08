package route

import (
	"template/domain/shared"
	valueobjects "template/domain/value-objects"
)

type WaypointType string

const (
	GreatCircle WaypointType = "GC"
	RhumbLine   WaypointType = "RL"
)

type Waypoint struct {
	id                valueobjects.WaypointID
	routeID           valueobjects.RouteID
	latitude          float64
	longitude         float64
	waypointID        *string
	orderNumber       int
	name              *string
	waypointType      WaypointType
	plannedSpeedKnots *float64
}

func (w Waypoint) ID() valueobjects.WaypointID {
	return w.id
}

func (w Waypoint) RouteID() valueobjects.RouteID {
	return w.routeID
}

func (w Waypoint) Latitude() float64 {
	return w.latitude
}

func (w Waypoint) Longitude() float64 {
	return w.longitude
}

func (w Waypoint) WaypointID() *string {
	return shared.CloneOptionalString(w.waypointID)
}

func (w Waypoint) OrderNumber() int {
	return w.orderNumber
}

func (w Waypoint) Name() *string {
	return shared.CloneOptionalString(w.name)
}

func (w Waypoint) WaypointType() WaypointType {
	return w.waypointType
}

func (w Waypoint) PlannedSpeedKnots() *float64 {
	return shared.CloneOptionalFloat64(w.plannedSpeedKnots)
}

func NewWaypoint(
	id valueobjects.WaypointID,
	routeID valueobjects.RouteID,
	latitude float64,
	longitude float64,
	waypointID *string,
	orderNumber int,
	name *string,
	waypointType WaypointType,
	plannedSpeedKnots *float64,
) (Waypoint, error) {
	if id.IsZero() {
		return Waypoint{}, ErrInvalidWaypointID
	}
	if routeID.IsZero() {
		return Waypoint{}, ErrInvalidRouteID
	}
	if !shared.IsValidCoordinates(latitude, longitude) {
		return Waypoint{}, ErrInvalidCoordinates
	}
	if orderNumber < 0 {
		return Waypoint{}, ErrInvalidOrderNumber
	}
	if !isValidWaypointType(waypointType) {
		return Waypoint{}, ErrInvalidWaypointType
	}

	return Waypoint{
		id:                id,
		routeID:           routeID,
		latitude:          latitude,
		longitude:         longitude,
		waypointID:        shared.CloneOptionalString(waypointID),
		orderNumber:       orderNumber,
		name:              shared.CloneOptionalString(name),
		waypointType:      waypointType,
		plannedSpeedKnots: shared.CloneOptionalFloat64(plannedSpeedKnots),
	}, nil
}

func ReconstructWaypoint(
	id valueobjects.WaypointID,
	routeID valueobjects.RouteID,
	latitude float64,
	longitude float64,
	waypointID *string,
	orderNumber int,
	name *string,
	waypointType WaypointType,
	plannedSpeedKnots *float64,
) (Waypoint, error) {
	return NewWaypoint(
		id,
		routeID,
		latitude,
		longitude,
		waypointID,
		orderNumber,
		name,
		waypointType,
		plannedSpeedKnots,
	)
}
