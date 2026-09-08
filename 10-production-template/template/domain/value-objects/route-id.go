package valueobjects

import "template/domain/domainerror"

type RouteID struct {
	value string
}

func NewRouteID(value string) (RouteID, error) {
	if err := parseUUID(value); err != nil {
		return RouteID{}, domainerror.New(domainerror.InvalidInput, "invalid route id")
	}
	return RouteID{value: value}, nil
}

func RouteIDFromTrusted(value string) RouteID {
	return RouteID{value: value}
}

func (id RouteID) String() string { return id.value }
func (id RouteID) IsZero() bool   { return id.value == "" }
