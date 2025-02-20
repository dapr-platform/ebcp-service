package api

import "github.com/go-chi/chi/v5"

func InitRoute(r chi.Router) {
	InitEbcp_exhibition_areaRoute(r)
	InitEbcp_exhibition_roomRoute(r)
	InitEbcp_exhibition_itemRoute(r)
	InitEbcp_exhibition_area_detailsRoute(r)
	InitEbcp_exhibition_hall_detailsRoute(r)
	InitEbcp_exhibition_hallRoute(r)
	InitEbcp_cameraRoute(r)
	InitEbcp_playerRoute(r)
	InitEbcp_playerExtRoute(r)
	InitEbcp_item_device_relationRoute(r)
	InitEbcp_item_scheduleRoute(r)
}
