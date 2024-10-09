package api

import "github.com/go-chi/chi/v5"

func InitRoute(r chi.Router) {
	InitEbcp_schedule_timeRoute(r)
	InitEbcp_exhibition_areaRoute(r)
	InitEbcp_exhibition_roomRoute(r)
	InitEbcp_exhibition_itemRoute(r)
	InitEbcp_deviceRoute(r)
	InitEbcp_exhibition_area_detailsRoute(r)
	InitEbcp_exhibition_hall_detailsRoute(r)
	InitEbcp_exhibition_hallRoute(r)
	InitEbcp_schedule_actionRoute(r)
	InitEbcp_schedule_taskRoute(r)
	InitEbcp_schedule_timeRoute(r)
}
