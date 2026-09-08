package valueobjects

import "template/domain/domainerror"

type WaypointID struct {
	value string
}

func NewWaypointID(value string) (WaypointID, error) {
	if err := parseUUID(value); err != nil {
		return WaypointID{}, domainerror.New(domainerror.InvalidInput, "invalid waypoint id")
	}
	return WaypointID{value: value}, nil
}

func WaypointIDFromTrusted(value string) WaypointID {
	return WaypointID{value: value}
}

func (id WaypointID) String() string { return id.value }
func (id WaypointID) IsZero() bool   { return id.value == "" }
