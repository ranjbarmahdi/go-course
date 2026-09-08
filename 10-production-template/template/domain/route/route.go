package route

import (
	"template/domain/shared"
	valueobjects "template/domain/value-objects"
	"time"
)

type RouteFileType string

const (
	NAVI_PLANNER_4000    RouteFileType = "navi-planner-4000"
	SIMPLE_TSV_WAYPOINTS RouteFileType = "simple-tsv-waypoints"
	NAVI_PLANNER_CSV     RouteFileType = "navi-planner-csv"
	VP_ROUTE             RouteFileType = "vp-route"
	RTZ                  RouteFileType = "rtz"
	RTZP                 RouteFileType = "rtzp"
	RTE                  RouteFileType = "rte"
)

type RouteCreatorType string

const (
	ECDIS  RouteCreatorType = "ecdis"
	User   RouteCreatorType = "user"
	VESSEL RouteCreatorType = "vessel"
)

type Route struct {
	id          valueobjects.RouteID
	vesselID    string
	name        string
	from        *string
	to          *string
	creatorType RouteCreatorType
	creatorID   *string
	createdAt   time.Time
	fileType    *RouteFileType
	waypoints   []Waypoint
}

func (r Route) ID() valueobjects.RouteID {
	return r.id
}

func (r Route) VesselID() string {
	return r.vesselID
}

func (r Route) Name() string {
	return r.name
}

func (r Route) From() *string {
	return shared.CloneOptionalString(r.from)
}

func (r Route) To() *string {
	return shared.CloneOptionalString(r.to)
}

func (r Route) CreatorType() RouteCreatorType {
	return r.creatorType
}

func (r Route) CreatorID() *string {
	return shared.CloneOptionalString(r.creatorID)
}

func (r Route) CreatedAt() time.Time {
	return r.createdAt
}

func (r Route) FileType() *RouteFileType {
	return shared.ClonePtr(r.fileType)
}

func (r Route) Waypoints() []Waypoint {
	out := make([]Waypoint, len(r.waypoints))
	copy(out, r.waypoints)
	return out
}

func NewRoute(
	id valueobjects.RouteID,
	vesselID string,
	name string,
	from *string,
	to *string,
	creatorType RouteCreatorType,
	creatorID *string,
	createdAt time.Time,
	fileType *RouteFileType,
) (*Route, error) {
	if id.IsZero() {
		return nil, ErrInvalidRouteID
	}
	if vesselID == "" {
		return nil, ErrInvalidVesselID
	}
	if name == "" {
		return nil, ErrInvalidName
	}
	if !isValidCreatorType(creatorType) {
		return nil, ErrInvalidCreatorType
	}

	normalizedCreatorID := shared.CloneOptionalString(creatorID)
	switch creatorType {
	case User:
		if normalizedCreatorID == nil {
			return nil, ErrInvalidCreatorID
		}
	default:
		normalizedCreatorID = nil
	}

	normalizedFileType := cloneOptionalRouteFileType(fileType)
	if normalizedFileType != nil && !isValidRouteFileType(*normalizedFileType) {
		return nil, ErrInvalidFileType
	}

	if createdAt.IsZero() {
		return nil, ErrInvalidCreatedAt
	}

	return &Route{
		id:          id,
		vesselID:    vesselID,
		name:        name,
		from:        shared.CloneOptionalString(from),
		to:          shared.CloneOptionalString(to),
		creatorType: creatorType,
		creatorID:   normalizedCreatorID,
		createdAt:   createdAt,
		fileType:    normalizedFileType,
		waypoints:   []Waypoint{},
	}, nil
}

func ReconstructRoute(
	id valueobjects.RouteID,
	vesselID string,
	name string,
	from *string,
	to *string,
	creatorType RouteCreatorType,
	creatorID *string,
	createdAt time.Time,
	fileType *RouteFileType,
	waypoints []Waypoint,
) (Route, error) {
	if id.IsZero() {
		return Route{}, ErrInvalidRouteID
	}

	return Route{
		id:          id,
		vesselID:    vesselID,
		name:        name,
		from:        shared.CloneOptionalString(from),
		to:          shared.CloneOptionalString(to),
		creatorType: creatorType,
		creatorID:   shared.CloneOptionalString(creatorID),
		createdAt:   createdAt,
		fileType:    cloneOptionalRouteFileType(fileType),
		waypoints:   append([]Waypoint(nil), waypoints...),
	}, nil
}

func (r *Route) AddWaypoint(wp Waypoint) error {
	if wp.RouteID() != r.id {
		return ErrWaypointRouteMismatch
	}

	for _, existing := range r.waypoints {
		if existing.OrderNumber() == wp.OrderNumber() {
			return ErrDuplicateOrderNumber
		}
	}

	r.waypoints = append(r.waypoints, wp)
	return nil
}
