package route

func cloneOptionalRouteFileType(v *RouteFileType) *RouteFileType {
	if v == nil {
		return nil
	}
	t := *v
	return &t
}

func isValidRouteFileType(t RouteFileType) bool {
	switch t {
	case NAVI_PLANNER_4000, SIMPLE_TSV_WAYPOINTS, NAVI_PLANNER_CSV,
		VP_ROUTE, RTZ, RTZP, RTE:
		return true
	default:
		return false
	}
}

func isValidCreatorType(t RouteCreatorType) bool {
	switch t {
	case ECDIS, User, VESSEL:
		return true
	default:
		return false
	}
}

func isValidWaypointType(t WaypointType) bool {
	return t == GreatCircle || t == RhumbLine
}
